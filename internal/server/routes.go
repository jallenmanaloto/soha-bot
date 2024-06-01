package server

import (
	"encoding/json"
	"net/http"

	"github.com/jallenmanaloto/soha-bot/pkg/logger"
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
		logger.Log.Errorf("ERROR unable to marshal health status: %v\n", err)
	}

	_, _ = w.Write(res)
}
