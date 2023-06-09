package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"go.dfds.cloud/aws-sso-exporter/aws"
	aConf "go.dfds.cloud/aws-sso-exporter/conf"
	awsssoexporter "go.dfds.cloud/aws-sso-exporter/metrics"
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
	respFed, err := aws.GetSsoEvents(cloudTrailClient, "ConsoleLogin")
	if err != nil {
		panic(err)
	}
	requestFinishedTime := time.Now()

	fmt.Printf("Request duration: %f\n", requestFinishedTime.Sub(requestStartTime).Seconds())
	fmt.Printf("Events retrieved: %d\n", len(respFed))

	federateStats := awsssoexporter.CalcConsoleLoginStats(respFed)
	fmt.Println("Console logins")
	fmt.Println("  Unique users: ", federateStats.UniqueUsers)
	fmt.Println("  Total logins: ", federateStats.TotalLogins)
	fmt.Println("  Average user logins: ", federateStats.AverageUserLogins)

	fmt.Println("Users:")
	for k, v := range federateStats.LoginsByUser {
		fmt.Printf("  %s: %d\n", k, v)
	}

}
