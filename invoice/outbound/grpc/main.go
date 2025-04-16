package grpc_outbound

import (
	"github.com/syauqeesy/accounting-service/invoice/configuration"
	"github.com/syauqeesy/accounting-service/proto/compiled/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCOutboundConnection struct {
	UserService *grpc.ClientConn
}

type GRPCOutboundService struct {
	Account account.AccountServiceClient
}

func New(configuration *configuration.Configuration) (*GRPCOutboundConnection, *GRPCOutboundService) {
	userServiceConnection, err := grpc.NewClient(configuration.GRPC.Service.User, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	grpcOutboundConnection := &GRPCOutboundConnection{
		UserService: userServiceConnection,
	}

	grpcOutboundService := &GRPCOutboundService{
		Account: account.NewAccountServiceClient(userServiceConnection),
	}

	return grpcOutboundConnection, grpcOutboundService
}

func (o *GRPCOutboundConnection) Close() error {
	if o.UserService == nil {
		return nil
	}

	o.UserService.Close()

	return nil
}
