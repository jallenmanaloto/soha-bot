package server

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	port int
}

const (
	ServerPort = 8000
)

func NewServer() *http.Server {
	NewServer := &Server{
		port: ServerPort,
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
