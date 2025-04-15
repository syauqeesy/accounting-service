package handler

import (
	"fmt"
	"net/http"
)

type accountHandler handler

func (h *accountHandler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "account register")
}
