package controllers

import (
	use_cases "api-bff-golang/domain/use-cases"
	"api-bff-golang/infrastructure/database/mongo/drivers/models"
)

type FindAllTaskControllerInterface interface {
	Process() ([]models.TaskMongoModel, error)
}

type FindAllTaskController struct {
	useCase use_cases.FindAllTaskUseCaseInterface
}

func NewFindAllTaskController(useCase use_cases.FindAllTaskUseCaseInterface) *FindAllTaskController {
	return &FindAllTaskController{useCase}
}

func (f *FindAllTaskController) Process() ([]models.TaskMongoModel, error) {
	return f.useCase.Process()
}
