package inputs

import (
	"api-bff-golang/infraestructure/web/models/response"
	"api-bff-golang/interfaces/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type DeleteTaskByTitleInputInterface interface {
	FromApi(c echo.Context) error
}
type DeleteTaskByTitleInput struct {
	controller controllers.DeleteTaskByTitleControllerInterface
}

func NewDeleteTaskByTitleInput(controller controllers.DeleteTaskByTitleControllerInterface) *DeleteTaskByTitleInput {
	return &DeleteTaskByTitleInput{controller}
}

func (d *DeleteTaskByTitleInput) FromApi(c echo.Context) error {
	r, e := d.controller.Process(c.Param("title"))

	if e != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Kind:        "SOME_ERROR",
			Description: "nothing to delete with title " + c.Param("title"),
		})
	}

	if r == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Kind:        "SOME_ERROR",
			Description: "nothing to delete with title " + c.Param("title"),
		})
	}
	s := strconv.FormatInt(r, 10)
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Data:        r,
		Description: "task deleted row affected  " + string(s),
	})

}
