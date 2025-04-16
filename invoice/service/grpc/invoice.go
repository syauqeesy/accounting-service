package grpc_service

import (
	"context"

	"github.com/syauqeesy/accounting-service/invoice/model"
	"github.com/syauqeesy/accounting-service/proto/compiled/invoice"
)

type invoiceService struct {
	grpcService
	invoice.UnimplementedInvoiceServiceServer
}

func (s *invoiceService) GetInvoiceInfo(invoiceModel *model.Invoice) *invoice.InvoiceInfo {
	return &invoice.InvoiceInfo{
		Id:        invoiceModel.Id,
		Email:     invoiceModel.Email,
		Amount:    invoiceModel.Amount,
		CreatedAt: invoiceModel.CreatedAt,
	}
}

func (s *invoiceService) SelectByUserId(ctx context.Context, request *invoice.SelectByUserIdRequest) (*invoice.SelectByUserIdResponse, error) {
	invoices, err := s.Repository.Invoice.SelectByUserId(request.GetUserId())
	if err != nil {
		return nil, err
	}

	invoiceInfos := make([]*invoice.InvoiceInfo, 0)

	for _, invoice := range invoices {
		invoiceInfos = append(invoiceInfos, s.GetInvoiceInfo(invoice))
	}

	return &invoice.SelectByUserIdResponse{
		Invoices: invoiceInfos,
	}, nil
}
