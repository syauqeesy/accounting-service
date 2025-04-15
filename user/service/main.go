package service

import "github.com/syauqeesy/accounting-service/configuration"

type service struct {
	Configuration *configuration.Configuration
}

type Service struct {
	Account AccountService
}

func New(configuration *configuration.Configuration) *Service {
	svc := &service{
		Configuration: configuration,
	}

	return &Service{
		Account: (*accountService)(svc),
	}
}
