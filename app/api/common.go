package api

import (
	"net/http"
	"encoding/json"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func ServeAsJSON(resp APIResponse, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(resp)
}
