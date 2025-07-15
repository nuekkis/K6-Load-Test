package handler

import (
    "encoding/json"
    "net/http"
    "time"
)

type HealthResponse struct {
    Status    string `json:"status"`
    Timestamp string `json:"timestamp"`
}

func Health(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    resp := HealthResponse{
        Status:    "ok",
        Timestamp: time.Now().Format(time.RFC3339),
    }
    json.NewEncoder(w).Encode(resp)
}
