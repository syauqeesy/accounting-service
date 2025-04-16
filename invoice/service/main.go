package service

import (
	"github.com/syauqeesy/accounting-service/invoice/configuration"
	grpc_outbound "github.com/syauqeesy/accounting-service/invoice/outbound/grpc"
	"github.com/syauqeesy/accounting-service/invoice/repository"
)

type service struct {
	Configuration *configuration.Configuration
	Repository    *repository.Repository
	GRPCOutbound  *grpc_outbound.GRPCOutboundService
}

type Service struct {
	Invoice InvoiceService
}

func New(configuration *configuration.Configuration, repository *repository.Repository, grpcOutbound *grpc_outbound.GRPCOutboundService) *Service {
	svc := &service{
		Configuration: configuration,
		Repository:    repository,
		GRPCOutbound:  grpcOutbound,
	}

	return &Service{
		Invoice: (*invoiceService)(svc),
	}
}
