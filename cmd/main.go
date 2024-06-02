package main

import (
	"github.com/jallenmanaloto/soha-bot/config"
	// "github.com/jallenmanaloto/soha-bot/internal/database"
	"github.com/jallenmanaloto/soha-bot/internal/server"
	"github.com/jallenmanaloto/soha-bot/pkg/logger"
)

func main() {
	// initialize aws config and load env variables
	config.NewClientConfig()

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
