package controllers

import (
	"api-bff-golang/domain/entities"
	use_cases "api-bff-golang/domain/use-cases"
	"api-bff-golang/infraestructure/database/mongo/drivers/models"
	log "api-bff-golang/infraestructure/logger"
)

type AddTaskControllerInterface interface {
	Process(t *entities.TaskEntity) (models.TaskMongoModel, error)
}

type AddTaskController struct {
	usecase use_cases.AddTaskUseCaseInterface
}

func NewAddTaskController(usecase use_cases.AddTaskUseCaseInterface) *AddTaskController {
	return &AddTaskController{usecase}
}

func (a *AddTaskController) Process(t *entities.TaskEntity) (models.TaskMongoModel, error) {
	log.Info("reading message in controller")
	return a.usecase.Process(t)
}
