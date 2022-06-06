package inputs

import (
	"api-bff-golang/infrastructure/web/models/response"
	"api-bff-golang/interfaces/controllers"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetTaskAndSendInputInterface interface {
	FromApi(ctx echo.Context) error
}

type GetTaskAndSendInput struct {
	controller controllers.GetTaskAndSendControllerInterface
}

func NewGetLastTaskAndSendInput(controller controllers.GetTaskAndSendControllerInterface) *GetTaskAndSendInput {
	return &GetTaskAndSendInput{controller}
}

func (g *GetTaskAndSendInput) FromApi(ctx echo.Context) error {
	e := g.controller.Process(ctx.Param("title"))
	if e != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Kind:        "SOME_ERROR",
			Description: fmt.Sprintf("failed with error %s", e.Error()),
		})
	}
	return ctx.JSON(http.StatusBadRequest, response.SuccessResponse{
		Data:        fmt.Sprintf("task sent %s", ctx.Param("title")),
		Description: fmt.Sprintf("Task send sucessfully "),
	})
}
