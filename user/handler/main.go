package handler

import (
	"net/http"

	"github.com/syauqeesy/accounting-service/user/configuration"
	"github.com/syauqeesy/accounting-service/user/service"
)

type handler struct {
	Service       *service.Service
	Configuration *configuration.Configuration
}

type Handler struct {
	Account *accountHandler
}

func New(mux *http.ServeMux, configuration *configuration.Configuration, service *service.Service) *Handler {
	handler := &handler{
		Configuration: configuration,
		Service:       service,
	}

	h := &Handler{
		Account: (*accountHandler)(handler),
	}

	mux.HandleFunc("GET /account", h.Account.List)

	return h
}
