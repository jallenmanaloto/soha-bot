package main

import (
	"github.com/jallenmanaloto/soha-bot/config"
	"github.com/jallenmanaloto/soha-bot/internal/bot"
	"github.com/jallenmanaloto/soha-bot/internal/database"
	"github.com/jallenmanaloto/soha-bot/internal/server"
	"github.com/jallenmanaloto/soha-bot/pkg/logger"
)

func main() {
	logger.Log.Info("INFO starting Application")

	// initialize aws config and load env variables
	config.NewClientConfig()

	// starting new db instance
	db := database.New()

	// initialize bot session
	logger.Log.Info("INFO starting Soha discord bot")
	bot, err := bot.New(db)
	if err != nil {
		logger.Log.Errorf("ERROR failed to initialize discord bot: %v", err)
	}

	err = bot.Session.Open()
	if err != nil {
		logger.Log.Errorf("ERROR failed to open discord bot session: %v", err)
	}

	// starting server for http reqs
	logger.Log.Info("INFO starting http server")
	server := server.NewServer()
	err = server.ListenAndServe()
	if err != nil {
		logger.Log.Errorf("ERROR unable to start server: %v\n", err)
		return
	}
}
