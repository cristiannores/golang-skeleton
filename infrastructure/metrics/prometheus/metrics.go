package prometheus

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type MetricPrometheus struct {
	Metric *prometheus.CounterVec
	Name   MetricNames
}

const APP = "golang-skeleton"

type MetricNames string

const (
	TEST  MetricNames = "test_metrics"
	ERROR MetricNames = "error_metrics"
)

var MetricsList = []MetricPrometheus{

	{
		Metric: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_%s", APP, ERROR),
			Help: "Total of times called this test metric",
		}, []string{"kind"}),
		Name: ERROR,
	},
}
