package main

import (
	"github.com/jallenmanaloto/soha-bot/config"
	// "github.com/jallenmanaloto/soha-bot/internal/database"
	"github.com/jallenmanaloto/soha-bot/internal/server"
	"github.com/jallenmanaloto/soha-bot/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	// initialize aws config and load env variables
	config.NewClientConfig()
	config.LoadEnvironmentVariables()

	if err := godotenv.Load(); err != nil {
		logger.Log.Errorf("ERROR unable to load env variables: %v\n", err)
		return
	}
	// db := database.New()

	logger.Log.Info("INFO starting Application")
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		logger.Log.Errorf("ERROR unable to start server: %v\n", err)
		return
	}

	logger.Log.Info("INFO server started...")
}
