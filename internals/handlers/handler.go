package handlers

import (
	"encoding/json"
	"net/http"
)

type APIHandler struct{}

func NewAPIHandler() *APIHandler {
    return &APIHandler{}
}

func (h *APIHandler) GetHello(w http.ResponseWriter, r *http.Request) {
    resp := map[string]string{"message": "Hello, World!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func (h *APIHandler) GetIndex(w http.ResponseWriter, r *http.Request) {
    resp := map[string]string{"status": "API is running"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}