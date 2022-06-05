package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type readinessHandler struct {
}

func NewReadinessHandler(e *echo.Echo) {
	h := &readinessHandler{}
	e.GET("/readiness", h.Readiness)
}

func (h *readinessHandler) Readiness(c echo.Context) error {
	//TODO: Pending validation with dependencies
	return c.JSON(http.StatusOK, "Readiness ok")
}
