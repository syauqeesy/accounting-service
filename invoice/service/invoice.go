package service

import "github.com/syauqeesy/accounting-service/invoice/payload"

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
		invoiceInfos = append(invoiceInfos, invoice.GetInfo())
	}

	return invoiceInfos, nil
}
