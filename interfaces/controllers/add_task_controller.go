package controllers

import (
	"api-bff-golang/domain/entities"
	use_cases "api-bff-golang/domain/use-cases"
	log "api-bff-golang/infrastructure/logger"
)

type AddTaskControllerInterface interface {
	Process(t *entities.TaskEntity) (entities.TaskEntity, error)
}

type AddTaskController struct {
	usecase use_cases.AddTaskUseCaseInterface
}

func NewAddTaskController(usecase use_cases.AddTaskUseCaseInterface) *AddTaskController {
	return &AddTaskController{usecase}
}

func (a *AddTaskController) Process(t *entities.TaskEntity) (entities.TaskEntity, error) {
	log.Info("reading message in controller")
	return a.usecase.Process(t)
}
