package main

import (
	"fmt"

	// "github.com/jallenmanaloto/soha-bot/internal/database"
	"github.com/jallenmanaloto/soha-bot/internal/server"
	"github.com/jallenmanaloto/soha-bot/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
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
