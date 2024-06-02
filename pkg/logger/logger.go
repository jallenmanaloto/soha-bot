package logger

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type LokiHook struct {
	Url string
}

func NewLokiHook(url string) *LokiHook {
	return &LokiHook{
		Url: url,
	}
}

func (h *LokiHook) Fire(entry *logrus.Entry) error {
	logMessage := [][]interface{}{
		{strconv.FormatInt(time.Now().UnixNano(), 10), entry.Message},
	}

	payload := map[string]interface{}{
		"streams": []map[string]interface{}{
			{
				"stream": map[string]interface{}{
					"Log": "value",
				},
				"values": logMessage,
			},
		},
	}
	jsonLog, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	username := os.Getenv("LOKI_USERNAME")
	password := os.Getenv("LOKI_PASSWORD")
	req, err := http.NewRequest("POST", h.Url, bytes.NewBuffer(jsonLog))
	if err != nil {
		return err
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func (h *LokiHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

var Log *logrus.Logger

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error: unable to load env variables")
	}

	Log = logrus.New()
	Log.SetLevel(logrus.InfoLevel)
	lokiLogUrl := os.Getenv("LOKI_LOG_URL")

	lokiHook := NewLokiHook(lokiLogUrl)
	Log.AddHook(lokiHook)
}
