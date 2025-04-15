package handler

import (
	"net/http"

	common_http "github.com/syauqeesy/accounting-service/common/http"
)

type accountHandler handler

func (h *accountHandler) Register(w http.ResponseWriter, r *http.Request) {
	common_http.WriteHttpResponse(w, http.StatusOK, http.StatusText(http.StatusOK), nil)
}
