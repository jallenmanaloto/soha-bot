package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.Health)

	return mux
}

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	healthStat := map[string]string{
		"status": "healthy",
		"db":     "connected",
	}
	res, err := json.Marshal(healthStat)
	if err != nil {
		log.Fatalf("Error handling marshal of health status: %v", err)
	}

	_, _ = w.Write(res)
}
