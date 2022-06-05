package web

import (
	log "api-bff-golang/infraestructure/logger"
	"api-bff-golang/infraestructure/web/routes"
	"api-bff-golang/interfaces/inputs"
	"api-bff-golang/shared/utils/config"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

var echoServer *echo.Echo

func NewWebServer() {
	echoServer = echo.New()
	echoServer.HideBanner = true
}

func InitRoutes(
	addTaskInput inputs.AddTaskInputInterface,
	getTaskByTitleInput inputs.GetTaskByTitleInputInterface,
	findAllTask inputs.FindAllTaskInputInterface,
	deleteTaskByTitle inputs.DeleteTaskByTitleInputInterface,
) {
	routes.NewHealthHandler(echoServer)
	routes.NewReadinessHandler(echoServer)
	routes.NewTaskHandler(echoServer, addTaskInput, getTaskByTitleInput, findAllTask, deleteTaskByTitle)
}

func Start() {
	log.Info("Config CurrentStage: %s ", config.GetString("currentStage"))
	log.Info("Config Developers: %v", config.GetArray("developers"))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.GetString("port")),
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}
	log.Info("App listen in port: %s", config.GetString("port"))
	echoServer.Logger.Fatal(echoServer.StartServer(server))
}

func Shutdown() {
	log.Info("Shutting down web server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := echoServer.Shutdown(ctx); err != nil {
		log.Fatal("Error shutting down web server")
	}
}
