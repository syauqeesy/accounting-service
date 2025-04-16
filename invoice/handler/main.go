package handler

import (
	"net/http"

	"github.com/syauqeesy/accounting-service/invoice/configuration"
	"github.com/syauqeesy/accounting-service/invoice/service"
)

type handler struct {
	Service       *service.Service
	Configuration *configuration.Configuration
}

type Handler struct {
	Invoice *invoiceHandler
}

func New(mux *http.ServeMux, configuration *configuration.Configuration, service *service.Service) *Handler {
	handler := &handler{
		Configuration: configuration,
		Service:       service,
	}

	h := &Handler{
		Invoice: (*invoiceHandler)(handler),
	}

	mux.HandleFunc("GET /invoice", h.Invoice.List)

	return h
}
