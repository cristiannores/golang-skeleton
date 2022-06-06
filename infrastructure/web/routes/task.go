package routes

import (
	"api-bff-golang/interfaces/inputs"
	"github.com/labstack/echo/v4"
)

type taskHandler struct {
	addTaskInput           inputs.AddTaskInputInterface
	getTaskByTitleInput    inputs.GetTaskByTitleInputInterface
	findAllTaskInput       inputs.FindAllTaskInputInterface
	deleteTaskByTitleInput inputs.DeleteTaskByTitleInputInterface
	getTaskAndSaveInput    inputs.GetTaskAndSendInputInterface
}

func NewTaskHandler(
	e *echo.Echo,
	addTaskInput inputs.AddTaskInputInterface,
	getTaskByTitleInput inputs.GetTaskByTitleInputInterface,
	findAllTaskInput inputs.FindAllTaskInputInterface,
	deleteTaskByTitleInput inputs.DeleteTaskByTitleInputInterface,
	getTaskAndSaveInput inputs.GetTaskAndSendInputInterface,
) {
	h := &taskHandler{
		addTaskInput,
		getTaskByTitleInput,
		findAllTaskInput,
		deleteTaskByTitleInput,
		getTaskAndSaveInput,
	}
	e.POST("/task", h.addTaskInput.FromApi)
	e.GET("/task/:title", h.getTaskByTitleInput.FromApi)
	e.GET("/task/send/:title", h.getTaskAndSaveInput.FromApi)
	e.GET("/task", h.findAllTaskInput.FromApi)
	e.DELETE("/task/:title", h.deleteTaskByTitleInput.FromApi)
}
