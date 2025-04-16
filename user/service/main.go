package service

import (
	"github.com/syauqeesy/accounting-service/user/configuration"
	grpc_outbound "github.com/syauqeesy/accounting-service/user/outbound/grpc"
	"github.com/syauqeesy/accounting-service/user/repository"
)

type service struct {
	Configuration *configuration.Configuration
	Repository    *repository.Repository
	GRPCOutbound  *grpc_outbound.GRPCOutboundService
}

type Service struct {
	Account AccountService
}

func New(configuration *configuration.Configuration, repository *repository.Repository, grpcOutbound *grpc_outbound.GRPCOutboundService) *Service {
	svc := &service{
		Configuration: configuration,
		Repository:    repository,
		GRPCOutbound:  grpcOutbound,
	}

	return &Service{
		Account: (*accountService)(svc),
	}
}
