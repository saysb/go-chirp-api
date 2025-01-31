// internal/handlers/response.go
package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
    Status  string      `json:"status"`
    Data    interface{} `json:"data,omitempty"`
    Message string      `json:"message,omitempty"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, statusCode int, message string) {
    JSON(w, statusCode, Response{
        Status:  "error",
        Message: message,
    })
}

func Success(w http.ResponseWriter, statusCode int, data interface{}) {
    JSON(w, statusCode, Response{
        Status: "success",
        Data:   data,
    })
}