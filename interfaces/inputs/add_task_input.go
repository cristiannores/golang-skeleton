package inputs

import (
	"api-bff-golang/domain/entities"
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/infrastructure/web/models/response"
	"api-bff-golang/interfaces/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AddTaskInputInterface interface {
	FromApi(c echo.Context) error
}

type AddTaskInput struct {
	controller controllers.AddTaskControllerInterface
}

func NewAddTaskInput(controller controllers.AddTaskControllerInterface) *AddTaskInput {
	return &AddTaskInput{controller}
}

func (t *AddTaskInput) FromApi(c echo.Context) error {

	log.Info("reading message in input")
	o := new(entities.TaskEntity)
	if err := c.Bind(o); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Kind:        "AN_ERROR",
			Description: err.Error(),
		})
	}
	r, e := t.controller.Process(o)

	if e != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Kind:        "AN_ERROR",
			Description: e.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		Data:        r,
		Description: "task saved successfully",
	})
}
