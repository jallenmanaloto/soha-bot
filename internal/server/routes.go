package server

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/jallenmanaloto/soha-bot/internal/constants"
	"github.com/jallenmanaloto/soha-bot/pkg/logger"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.Health)
	mux.HandleFunc("POST /webhook/alert", s.UpdateAlert)

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

type Alert struct {
	Signature string `json:"signature"`
	Title     string `json:"title"`
	ManhwaId  string `json:"manhwaId"`
}

type AlertResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (s *Server) UpdateAlert(w http.ResponseWriter, r *http.Request) {
	var alert Alert
	sigKey := os.Getenv("SCRAPER_KEY")
	err := json.NewDecoder(r.Body).Decode(&alert)
	if err != nil {
		logger.Log.Errorf(constants.ErrorJsonDecode, err)
	}

	if sigKey != alert.Signature {
		alertRes := &AlertResponse{
			Status:  401,
			Message: constants.Unauthorized,
		}
		res, err := json.Marshal(alertRes)
		if err != nil {
			logger.Log.Errorf(constants.ErrorMarshalRes, err)
		}
		_, _ = w.Write(res)
		return
	}

	// start operation
	// go routine to send discord mess & send json

	mes := &AlertResponse{
		Status:  200,
		Message: "Accepted",
	}
	succ, err := json.Marshal(mes)
	if err != nil {
		logger.Log.Errorf(constants.ErrorMarshalRes, err)
	}
	_, _ = w.Write(succ)
}
