package prometheus

import (
	log "api-bff-golang/infrastructure/logger"
	"github.com/prometheus/client_golang/prometheus"
)

type MetricInterface interface {
	InitMetrics()
	GetMetricByName(name MetricNames) *prometheus.CounterVec
	IncrementMetricError(labelValue string) error
}

type MetricClient struct {
	metrics []MetricPrometheus
}

func NewMetric() *MetricClient {
	metrics := MetricsList
	return &MetricClient{metrics}
}

func (m *MetricClient) InitMetrics() {
	for _, metric := range m.metrics {
		m.AddMetric(metric.Metric)
		log.Debug("metric initialized %s", metric.Name)
	}
}

func (m *MetricClient) AddMetric(metric *prometheus.CounterVec) {
	prometheus.Register(metric)
}

func (m *MetricClient) GetMetricByName(name MetricNames) *prometheus.CounterVec {
	log.Info("Getting metric by name %s", name)
	for _, metric := range m.metrics {
		if metric.Name == name {
			return metric.Metric
		}
	}

	log.Info("Error getting metric by name %s", name)
	return nil
}

func (m *MetricClient) IncrementMetricError(labelValue string) error {
	log.Info("increment metric error label: %s", labelValue)

	m.GetMetricByName(ERROR).With(prometheus.Labels{
		"kind": labelValue,
	}).Inc()
	return nil
}