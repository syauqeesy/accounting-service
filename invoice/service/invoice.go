package service

import (
	"context"

	"github.com/syauqeesy/accounting-service/invoice/payload"
	"github.com/syauqeesy/accounting-service/proto/compiled/account"
)

type InvoiceService interface {
	List() ([]*payload.InvoiceInfo, error)
}

type invoiceService service

func (s *invoiceService) List() ([]*payload.InvoiceInfo, error) {
	invoiceInfos := make([]*payload.InvoiceInfo, 0)

	invoices, err := s.Repository.Invoice.Select()
	if err != nil {
		return nil, err
	}

	for _, invoice := range invoices {
		result, err := s.GRPCOutbound.Account.SelectById(context.Background(), &account.SelectByIdRequest{
			Id: invoice.UserId,
		})
		if err != nil {
			return nil, err
		}

		invoiceInfos = append(invoiceInfos, invoice.GetInfo(result))
	}

	return invoiceInfos, nil
}
