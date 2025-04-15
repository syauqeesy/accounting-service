package http

import (
	"encoding/json"
	"net/http"
)

func HandleHttpError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	err := json.NewEncoder(w).Encode(&HttpJsonResponse{
		Message: http.StatusText(http.StatusInternalServerError),
		Data:    nil,
	})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
