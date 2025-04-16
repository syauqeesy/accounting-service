package grpc_outbound

import (
	"github.com/syauqeesy/accounting-service/proto/compiled/invoice"
	"github.com/syauqeesy/accounting-service/user/configuration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCOutboundConnection struct {
	InvoiceService *grpc.ClientConn
}

type GRPCOutboundService struct {
	Invoice invoice.InvoiceServiceClient
}

func New(configuration *configuration.Configuration) (*GRPCOutboundConnection, *GRPCOutboundService) {
	invoiceServiceConnection, err := grpc.NewClient(configuration.GRPC.Service.Invoice, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	grpcOutboundConnection := &GRPCOutboundConnection{
		InvoiceService: invoiceServiceConnection,
	}

	grpcOutboundService := &GRPCOutboundService{
		Invoice: invoice.NewInvoiceServiceClient(invoiceServiceConnection),
	}

	return grpcOutboundConnection, grpcOutboundService
}

func (o *GRPCOutboundConnection) Close() error {
	if o.InvoiceService == nil {
		return nil
	}

	o.InvoiceService.Close()

	return nil
}
