package service

import (
	"github.com/syauqeesy/accounting-service/user/configuration"
	"github.com/syauqeesy/accounting-service/user/repository"
)

type service struct {
	Configuration *configuration.Configuration
	Repository    *repository.Repository
}

type Service struct {
	Account AccountService
}

func New(configuration *configuration.Configuration, repository *repository.Repository) *Service {
	svc := &service{
		Configuration: configuration,
		Repository:    repository,
	}

	return &Service{
		Account: (*accountService)(svc),
	}
}
