package service

import (
	"github.com/syauqeesy/accounting-service/invoice/configuration"
	"github.com/syauqeesy/accounting-service/invoice/repository"
)

type service struct {
	Configuration *configuration.Configuration
	Repository    *repository.Repository
}

type Service struct {
	Invoice InvoiceService
}

func New(configuration *configuration.Configuration, repository *repository.Repository) *Service {
	svc := &service{
		Configuration: configuration,
		Repository:    repository,
	}

	return &Service{
		Invoice: (*invoiceService)(svc),
	}
}
