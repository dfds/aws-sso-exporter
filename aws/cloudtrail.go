package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsHttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"go.uber.org/ratelimit"
	"net/http"
	"time"
)

func CreateHttpClientWithoutKeepAlive() *awsHttp.BuildableClient {
	client := awsHttp.NewBuildableClient().WithTransportOptions(func(transport *http.Transport) {
		transport.DisableKeepAlives = true
	})

	return client
}

func GetSsoEvents(client *cloudtrail.Client, eventName string) ([]types.Event, error) {
	var results []types.Event

	startTime, err := time.Parse(time.DateOnly, "2023-02-01")
	if err != nil {
		return results, err
	}

	resp, err := client.LookupEvents(context.Background(), &cloudtrail.LookupEventsInput{
		EndTime: nil,
		LookupAttributes: []types.LookupAttribute{
			{
				AttributeKey:   "EventName",
				AttributeValue: aws.String(eventName),
			},
		},
		NextToken: nil,
		StartTime: aws.Time(startTime),
	})
	if err != nil {
		return results, err
	}

	results = append(results, resp.Events...)

	rl := ratelimit.New(1)
	for resp.NextToken != nil {
		rl.Take()
		resp, err = client.LookupEvents(context.Background(), &cloudtrail.LookupEventsInput{
			LookupAttributes: []types.LookupAttribute{
				{
					AttributeKey:   "EventName",
					AttributeValue: aws.String(eventName),
				},
			},
			NextToken: resp.NextToken,
			StartTime: aws.Time(startTime),
		})

		if err != nil {
			return results, err
		}
		results = append(results, resp.Events...)

	}

	return results, nil
}

func GetSsoFederateEvent(client *cloudtrail.Client) ([]types.Event, error) {
	return GetSsoEvents(client, "Federate")
}

func GetSsoAuthenticateEvent(client *cloudtrail.Client) ([]types.Event, error) {
	return GetSsoEvents(client, "Authenticate")
}
