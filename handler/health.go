package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

var startTime = time.Now()

type HealthResponse struct {
	Status  string  `json:"status"`
	Model   string  `json:"model"`
	Version string  `json:"version"`
	Uptime  float64 `json:"uptime_seconds"`
}

func Health(cfg interface{ GetModel() string }) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(HealthResponse{
			Status:  "ok",
			Model:   "whisper-medium-am-v1-47wer-v2",
			Version: "1.0.0",
			Uptime:  time.Since(startTime).Seconds(),
		})
	}
}
