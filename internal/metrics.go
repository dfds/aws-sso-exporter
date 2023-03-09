package internal

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var FederateTotalUniqueUsersGauge prometheus.Gauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name:      "unique_users",
	Help:      "Total unique users",
	Subsystem: "federate",
	Namespace: "aws_sso_exporter",
})

var FederateTotalRolesAssumedGauge prometheus.Gauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name:      "roles_assumed",
	Help:      "Total roles assumed",
	Subsystem: "federate",
	Namespace: "aws_sso_exporter",
})

var FederateAverageRolesAssumedGauge prometheus.Gauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name:      "average_roles_assumed",
	Help:      "Average roles assumed",
	Subsystem: "federate",
	Namespace: "aws_sso_exporter",
})

var FederateMeanRolesAssumedGauge prometheus.Gauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name:      "mean_roles_assumed",
	Help:      "Mean roles assumed",
	Subsystem: "federate",
	Namespace: "aws_sso_exporter",
})

var AuthenticateTotalUniqueUsersGauge prometheus.Gauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name:      "unique_users",
	Help:      "Total unique users",
	Subsystem: "authenticate",
	Namespace: "aws_sso_exporter",
})

var AuthenticateTotalSignInsAssumedGauge prometheus.Gauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name:      "total_signins",
	Help:      "Total sign-ins",
	Subsystem: "authenticate",
	Namespace: "aws_sso_exporter",
})

var AuthenticateAverageUserSignInsAssumedGauge prometheus.Gauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name:      "average_user_signin",
	Help:      "Average user sign in",
	Subsystem: "authenticate",
	Namespace: "aws_sso_exporter",
})

var AuthenticateMeanUserSignInsAssumedGauge prometheus.Gauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name:      "mean_user_signin",
	Help:      "Mean user sign in",
	Subsystem: "authenticate",
	Namespace: "aws_sso_exporter",
})
