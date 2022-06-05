package main

import (
	"api-bff-golang/infraestructure"
	"api-bff-golang/infraestructure/database/mongo/client"
	log "api-bff-golang/infraestructure/logger"
	"api-bff-golang/infraestructure/web"
	"api-bff-golang/shared/utils/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	log.Info("starting golang app")
	config.LoadSettings()

	//Signs Catcher
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	web.NewWebServer()

	mongoClient := infraestructure.SetupDependencies()

	go web.Start()

	//Graceful Shutdown process
	sig := <-quit
	gracefulShutdown(sig, mongoClient)

}

func gracefulShutdown(sig os.Signal, mongoClient *client.MongoClient) {

	log.Info("Signal trap : %v\n", sig)
	web.Shutdown()
	mongoClient.Disconnect()
	log.Info("Shutdown process completed for signal: %v\n", sig)
}
