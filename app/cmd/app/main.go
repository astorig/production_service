package main

import (
	"log"
	"product-service/internal/app"

	//"product-service/internal/app"
	"product-service/internal/config"
	"product-service/pkg/logging"
)

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()

	log.Print("logger initializing")
	logging.Init()
	logger := logging.GetLogger()
	a, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Running Application")
	a.Run()
}
