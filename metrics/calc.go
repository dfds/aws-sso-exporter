package metrics

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"go.dfds.cloud/aws-sso-exporter/aws"
	"sort"
)

type StatsFederate struct {
	TotalRolesAssumed             int
	UniqueUsers                   int
	AverageUserRoleAssumed        int
	MeanUserRoleAssumed           float64
	RoleAssumedCountByUser        map[string]int
	RoleAssumedCountryTotalCount  map[string]int
	RoleAssumedCountryUniqueUsers map[string]map[string]int
}

type StatsAuthenticate struct {
	TotalSignins             int
	UniqueUsers              int
	AverageUserSignIn        int
	MeanUserSignIn           float64
	SignInCountByUser        map[string]int
	SignInCountryTotalCount  map[string]int
	SignInCountryUniqueUsers map[string]map[string]int
}

func CalcFederateStats(events []types.Event) StatsFederate {
	stats := StatsFederate{
		TotalRolesAssumed:             0,
		UniqueUsers:                   0,
		AverageUserRoleAssumed:        0,
		MeanUserRoleAssumed:           0,
		RoleAssumedCountByUser:        map[string]int{},
		RoleAssumedCountryTotalCount:  map[string]int{},
		RoleAssumedCountryUniqueUsers: map[string]map[string]int{},
	}

	for _, event := range events {
		stats.TotalRolesAssumed = stats.TotalRolesAssumed + 1
		stats.RoleAssumedCountByUser[*event.Username] = stats.RoleAssumedCountByUser[*event.Username] + 1

		//fmt.Printf("%s, %s, %s\n", *event.EventName, *event.Username, *event.EventTime)
		var innerEvent aws.FederateEvent
		err := json.Unmarshal([]byte(*event.CloudTrailEvent), &innerEvent)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("  %s - %s accessing %s in %s\n", innerEvent.AwsRegion, innerEvent.UserIdentity.UserName, innerEvent.ServiceEventDetails.RoleName, innerEvent.ServiceEventDetails.AccountID)
	}

	stats.UniqueUsers = len(stats.RoleAssumedCountByUser)
	stats.AverageUserRoleAssumed = stats.TotalRolesAssumed / stats.UniqueUsers

	data := []float64{}
	for _, entry := range stats.RoleAssumedCountByUser {
		data = append(data, float64(entry))
	}

	stats.MeanUserRoleAssumed = calcMedian(data)

	return stats
}

func CalcAuthenticateStats(events []types.Event) StatsAuthenticate {
	stats := StatsAuthenticate{
		TotalSignins:             0,
		UniqueUsers:              0,
		AverageUserSignIn:        0,
		MeanUserSignIn:           0,
		SignInCountByUser:        map[string]int{},
		SignInCountryTotalCount:  map[string]int{},
		SignInCountryUniqueUsers: map[string]map[string]int{},
	}

	for _, event := range events {
		stats.TotalSignins = stats.TotalSignins + 1
		stats.SignInCountByUser[*event.Username] = stats.SignInCountByUser[*event.Username] + 1

		//fmt.Printf("%s, %s, %s\n", *event.EventName, *event.Username, *event.EventTime)
		var innerEvent aws.AuthenticateEvent
		err := json.Unmarshal([]byte(*event.CloudTrailEvent), &innerEvent)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("  %s - %s\n", innerEvent.AwsRegion, innerEvent.UserIdentity.UserName)
	}

	stats.UniqueUsers = len(stats.SignInCountByUser)
	stats.AverageUserSignIn = stats.TotalSignins / stats.UniqueUsers

	data := []float64{}
	for _, entry := range stats.SignInCountByUser {
		data = append(data, float64(entry))
	}

	stats.MeanUserSignIn = calcMedian(data)

	return stats
}

func calcMedian(n []float64) float64 {
	sort.Float64s(n) // sort the numbers

	mNumber := len(n) / 2

	if isOdd(n) {
		return n[mNumber]
	}

	return (n[mNumber-1] + n[mNumber]) / 2
}

func isOdd(n []float64) bool {
	if len(n)%2 == 0 {
		return false
	}

	return true
}
