package httpx

import (
	"encoding/json"
	"net/http"
	"seat-map/internal/dto"
)

func ResponseJson(w http.ResponseWriter, statusCode int, code string, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(dto.APIResponse{Code: code, Message: message, Data: data})
}
