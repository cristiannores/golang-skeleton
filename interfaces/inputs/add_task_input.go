package inputs

import (
	"api-bff-golang/domain/entities"
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/infrastructure/web/models/response"
	"api-bff-golang/interfaces/controllers"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type AddTaskInputInterface interface {
	FromApi(c echo.Context) error
	FromKafka(value []byte) error
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

func (t *AddTaskInput) FromKafka(value []byte) error {

	task := entities.TaskEntity{}

	if err := json.Unmarshal(value, &task); err != nil {
		log.Error("Error converting checkout order to ox order %s", err.Error())
		return nil
	}
	currentTime := time.Now()
	task.Tags = append(task.Tags, fmt.Sprintf("received in kafka  : %s", currentTime.Format("2006.01.02 15:04:05")))
	_, e := t.controller.Process(&task)

	if e != nil {
		log.Error("Error faving task from kafka error %s", e.Error())
		return e
	}

	log.Info("Task saved sucessfully from kafka")
	return nil
}
