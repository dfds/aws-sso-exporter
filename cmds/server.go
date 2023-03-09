package main

import (
	"context"
	"fmt"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.dfds.cloud/aws-sso-exporter/aws"
	"go.dfds.cloud/aws-sso-exporter/conf"
	"go.dfds.cloud/aws-sso-exporter/internal"
	"go.dfds.cloud/aws-sso-exporter/metrics"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(pprof.New())

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	go worker()
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}

func worker() {
	config, err := conf.LoadConfig()
	if err != nil {
		panic(err)
	}

	sleepInterval, err := time.ParseDuration(fmt.Sprintf("%ds", config.WorkerInterval))
	if err != nil {
		panic(err)
	}

	for {
		cfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithRegion(config.Aws.Region), awsConfig.WithHTTPClient(aws.CreateHttpClientWithoutKeepAlive()))
		if err != nil {
			panic(err)
		}

		cloudTrailClient := cloudtrail.NewFromConfig(cfg)

		requestStartTime := time.Now()
		respFed, err := aws.GetSsoFederateEvent(cloudTrailClient)
		if err != nil {
			panic(err)
		}
		requestFinishedTime := time.Now()

		fmt.Printf("Request duration: %f\n", requestFinishedTime.Sub(requestStartTime).Seconds())
		fmt.Printf("Events retrieved: %d\n", len(respFed))

		requestStartTime = time.Now()
		respAuth, err := aws.GetSsoAuthenticateEvent(cloudTrailClient)
		if err != nil {
			panic(err)
		}
		requestFinishedTime = time.Now()

		fmt.Printf("Request duration: %f\n", requestFinishedTime.Sub(requestStartTime).Seconds())
		fmt.Printf("Events retrieved: %d\n", len(respAuth))

		federateStats := metrics.CalcFederateStats(respFed)
		authenticateStats := metrics.CalcAuthenticateStats(respAuth)

		// Federate
		internal.FederateTotalUniqueUsersGauge.Set(float64(federateStats.UniqueUsers))
		internal.FederateTotalRolesAssumedGauge.Set(float64(federateStats.TotalRolesAssumed))
		internal.FederateAverageRolesAssumedGauge.Set(float64(federateStats.AverageUserRoleAssumed))
		internal.FederateMeanRolesAssumedGauge.Set(federateStats.MeanUserRoleAssumed)

		// Authenticate
		internal.AuthenticateTotalUniqueUsersGauge.Set(float64(authenticateStats.UniqueUsers))
		internal.AuthenticateTotalSignInsAssumedGauge.Set(float64(authenticateStats.TotalSignins))
		internal.AuthenticateAverageUserSignInsAssumedGauge.Set(float64(authenticateStats.AverageUserSignIn))
		internal.AuthenticateMeanUserSignInsAssumedGauge.Set(authenticateStats.MeanUserSignIn)

		fmt.Println("New SSO metrics published")
		time.Sleep(time.Second * sleepInterval)
	}
}
