package main

import (
	"fmt"

	"github.com/jallenmanaloto/soha-bot/internal/database"
	"github.com/jallenmanaloto/soha-bot/internal/server"
)

func main() {
	db := database.New()

	server := server.NewServer()
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Unable to start server: %s", err))
	}
}
