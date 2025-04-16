package repository

import "github.com/syauqeesy/accounting-service/invoice/model"

type InvoiceRepository interface {
	Select() ([]*model.Invoice, error)
	SelectByUserId(userId string) ([]*model.Invoice, error)
}

type invoiceRepository repository

func (r *invoiceRepository) Select() ([]*model.Invoice, error) {
	invoices := make([]*model.Invoice, 0)

	q := r.Database.Find(&invoices)
	if q.Error != nil {
		return nil, q.Error
	}

	return invoices, nil
}

func (r *invoiceRepository) SelectByUserId(userId string) ([]*model.Invoice, error) {
	invoices := make([]*model.Invoice, 0)

	q := r.Database.Where("user_id = ?", userId).Find(&invoices)
	if q.Error != nil {
		return nil, q.Error
	}

	return invoices, nil
}
