package routes

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

type healthHandler struct {
}

type health struct {
	Status      string `json:"status"`
	Version     string `json:"version"`
	Uptime      string `json:"uptime"`
	Environment string `json:"environment"`
	Region      string `json:"region"`
}

func NewHealthHandler(e *echo.Echo) {
	h := &healthHandler{}
	e.GET("/health", h.HealthCheck)
}

func (p *healthHandler) HealthCheck(c echo.Context) error {
	versionApp := os.Getenv("VERSION_APP")
	healthCheck := health{
		Status:      "UP",
		Version:     versionApp,
		Uptime:      time.Since(startTime).String(),
		Environment: os.Getenv("ENVIRONMENT"),
		Region:      os.Getenv("REGION"),
	}
	return c.JSON(http.StatusOK, healthCheck)
}
