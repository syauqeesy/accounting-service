package http

import (
	"encoding/json"
	"net/http"
)

type HttpJsonResponse struct {
	Message string       `json:"message"`
	Data    *interface{} `json:"data"`
}

func WriteHttpResponse(w http.ResponseWriter, status int, message string, payload *interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(&HttpJsonResponse{
		Message: message,
		Data:    payload,
	})
	if err != nil {
		HandleHttpError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}
