package controllers

import (
	use_cases "api-bff-golang/domain/use-cases"
	"api-bff-golang/infraestructure/database/mongo/drivers/models"
)

type GetTaskByNameControllerInterface interface {
	Process(title string) (models.TaskMongoModel, error)
}

type GetTaskByNameController struct {
	usecase use_cases.GetTaskByNameUseCaseInterface
}

func NewGetTaskByNameController(usecase use_cases.GetTaskByNameUseCaseInterface) *GetTaskByNameController {
	return &GetTaskByNameController{usecase}
}

func (c GetTaskByNameController) Process(title string) (models.TaskMongoModel, error) {
	return c.usecase.Process(title)
}
