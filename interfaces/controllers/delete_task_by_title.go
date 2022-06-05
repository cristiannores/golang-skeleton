package controllers

import use_cases "api-bff-golang/domain/use-cases"

type DeleteTaskByTitleControllerInterface interface {
	Process(title string) (int64, error)
}

type DeleteTaskByTitleController struct {
	useCase use_cases.DeleteTaskByTitleUseCaseInterface
}

func NewDeleteTaskByTitleController(useCase use_cases.DeleteTaskByTitleUseCaseInterface) *DeleteTaskByTitleController {
	return &DeleteTaskByTitleController{useCase}
}

func (d *DeleteTaskByTitleController) Process(title string) (int64, error) {
	return d.useCase.Process(title)
}
