package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jallenmanaloto/soha-bot/internal/bot"
)

type Server struct {
	port int
	Bot  *bot.DiscordBot
}

const (
	ServerPort = 8000
)

func NewServer(bot *bot.DiscordBot) *http.Server {
	NewServer := &Server{
		port: ServerPort,
		Bot:  bot,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", ServerPort),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
