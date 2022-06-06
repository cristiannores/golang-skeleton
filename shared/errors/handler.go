package errors

import (
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/infrastructure/metrics/prometheus"
	"errors"
)

func ErrorHandler(e error, metricClient prometheus.MetricInterface) {
	log.Info("[ErrorHandler] handling message error")
	var customError *CustomError
	switch {
	case errors.As(e, &customError):
		log.Info("Incrementic metric on error handler")
		metricClient.IncrementMetricError(string(customError.Kind))
		log.Info("Incrementic  metric success on error handler")
	default:
		metricClient.IncrementMetricError(string(UNEXPECTED_ERROR))
	}
}

func GetKind(e error) string {
	var customError *CustomError
	switch {
	case errors.As(e, &customError):
		return string(customError.Kind)
	default:
		return string(UNEXPECTED_ERROR)
	}

}
