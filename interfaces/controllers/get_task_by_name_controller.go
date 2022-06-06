package controllers

import (
	use_cases "api-bff-golang/domain/use-cases"
	"api-bff-golang/infrastructure/database/mongo/drivers/models"
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/infrastructure/metrics/prometheus"
	"api-bff-golang/shared/errors"
)

type GetTaskByNameControllerInterface interface {
	Process(title string) (models.TaskMongoModel, error)
}

type GetTaskByNameController struct {
	usecase use_cases.GetTaskByNameUseCaseInterface
	metric  prometheus.MetricInterface
}

func NewGetTaskByNameController(usecase use_cases.GetTaskByNameUseCaseInterface, metric prometheus.MetricInterface) *GetTaskByNameController {
	return &GetTaskByNameController{usecase, metric}
}

func (c GetTaskByNameController) Process(title string) (models.TaskMongoModel, error) {
	log.Info("[get_task_by_name_controller] init controller with params %s", title)

	r, e := c.usecase.Process(title)

	if e != nil {
		errors.ErrorHandler(e, c.metric)
		return models.TaskMongoModel{}, e
	}
	return r, nil
}
