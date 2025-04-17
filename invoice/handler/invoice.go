package handler

import (
	"net/http"

	common_http "github.com/syauqeesy/accounting-service/common/http"
)

type invoiceHandler handler

func (h *invoiceHandler) List(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.Invoice.List()
	if err != nil {
		common_http.HandleHttpError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	common_http.WriteHttpResponse(w, http.StatusOK, http.StatusText(http.StatusOK), result)
}
