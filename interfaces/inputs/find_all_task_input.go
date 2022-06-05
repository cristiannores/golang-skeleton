package inputs

import (
	"api-bff-golang/infrastructure/web/models/response"
	"api-bff-golang/interfaces/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type FindAllTaskInputInterface interface {
	FromApi(c echo.Context) error
}

type FindAllTaskInput struct {
	controller controllers.FindAllTaskControllerInterface
}

func NewFindAllTaskInput(controller controllers.FindAllTaskControllerInterface) *FindAllTaskInput {
	return &FindAllTaskInput{controller}
}

func (f *FindAllTaskInput) FromApi(c echo.Context) error {

	r, e := f.controller.Process()

	if e != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Kind:        "SOME_ERROR",
			Description: "tasks not found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Data:        r,
		Description: "Tasks found successfully",
	})
}
