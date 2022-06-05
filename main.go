package main

import (
	"api-bff-golang/infrastructure"
	"api-bff-golang/infrastructure/database/mongo/client"
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/infrastructure/web"
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

	mongoClient := infrastructure.SetupDependencies()

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
