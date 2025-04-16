package grpc_service

import (
	"github.com/syauqeesy/accounting-service/invoice/configuration"
	"github.com/syauqeesy/accounting-service/invoice/repository"
	"github.com/syauqeesy/accounting-service/proto/compiled/invoice"
	"google.golang.org/grpc"
)

type grpcService struct {
	Configuration *configuration.Configuration
	Repository    *repository.Repository
}

func New(server *grpc.Server, configuration *configuration.Configuration, repository *repository.Repository) {
	svc := &grpcService{
		Configuration: configuration,
		Repository:    repository,
	}

	invoiceService := &invoiceService{
		grpcService: *svc,
	}

	invoice.RegisterInvoiceServiceServer(server, invoiceService)
}
