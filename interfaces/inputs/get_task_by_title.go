package inputs

import (
	"api-bff-golang/infraestructure/web/models/response"
	"api-bff-golang/interfaces/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetTaskByTitleInputInterface interface {
	FromApi(c echo.Context) error
}

type GetTaskByTitleInput struct {
	contoller controllers.GetTaskByNameControllerInterface
}

func NewGetTaskByTitleInput(contoller controllers.GetTaskByNameControllerInterface) *GetTaskByTitleInput {
	return &GetTaskByTitleInput{contoller}
}

func (g *GetTaskByTitleInput) FromApi(c echo.Context) error {
	r, e := g.contoller.Process(c.Param("title"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Kind:        "SOME_ERROR",
			Description: e.Error() + " " + c.Param("title"),
		})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		Data:        r,
		Description: "task found successfully",
	})
}
