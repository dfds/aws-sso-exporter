package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	awsssoexporter "go.dfds.cloud/aws-sso-exporter"
	"go.dfds.cloud/aws-sso-exporter/aws"
	aConf "go.dfds.cloud/aws-sso-exporter/conf"
	"time"
)

func main() {
	conf, err := aConf.LoadConfig()
	if err != nil {
		panic(err)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(conf.Aws.Region), config.WithHTTPClient(aws.CreateHttpClientWithoutKeepAlive()))
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
	fmt.Printf("Events retrieved: %d\n", len(respFed))

	federateStats := awsssoexporter.CalcFederateStats(respFed)
	authenticateStats := awsssoexporter.CalcAuthenticateStats(respAuth)

	fmt.Println("Federate")
	fmt.Println("  Unique users: ", federateStats.UniqueUsers)
	fmt.Println("  Total roles assumed: ", federateStats.TotalRolesAssumed)
	fmt.Println("  Average roles assumed: ", federateStats.AverageUserRoleAssumed)
	fmt.Println("  Mean roles assumed: ", federateStats.MeanUserRoleAssumed)

	fmt.Println("Authenticate")
	fmt.Println("  Unique users: ", authenticateStats.UniqueUsers)
	fmt.Println("  Total sign-ins: ", authenticateStats.TotalSignins)
	fmt.Println("  Average user sign in: ", authenticateStats.AverageUserSignIn)
	fmt.Println("  Mean user sign in: ", authenticateStats.MeanUserSignIn)
}
