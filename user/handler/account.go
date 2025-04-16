package handler

import (
	"net/http"

	common_http "github.com/syauqeesy/accounting-service/common/http"
)

type accountHandler handler

func (h *accountHandler) Register(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.Account.List()
	if err != nil {
		common_http.HandleHttpError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	common_http.WriteHttpResponse(w, http.StatusOK, http.StatusText(http.StatusOK), result)
}
