package controllers

import (
	"api-bff-golang/domain/entities"
	use_cases "api-bff-golang/domain/use-cases"
)

type FindAllTaskControllerInterface interface {
	Process() ([]entities.TaskEntity, error)
}

type FindAllTaskController struct {
	useCase use_cases.FindAllTaskUseCaseInterface
}

func NewFindAllTaskController(useCase use_cases.FindAllTaskUseCaseInterface) *FindAllTaskController {
	return &FindAllTaskController{useCase}
}

func (f *FindAllTaskController) Process() ([]entities.TaskEntity, error) {
	return f.useCase.Process()
}
