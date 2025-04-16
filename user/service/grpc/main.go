package grpc_service

import (
	"github.com/syauqeesy/accounting-service/proto/compiled/account"
	"github.com/syauqeesy/accounting-service/user/configuration"
	"github.com/syauqeesy/accounting-service/user/repository"
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

	accountService := &accountService{
		grpcService: *svc,
	}

	account.RegisterAccountServiceServer(server, accountService)
}
