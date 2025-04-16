package model

import (
	"github.com/syauqeesy/accounting-service/proto/compiled/invoice"
	"github.com/syauqeesy/accounting-service/user/payload"
)

type Account struct {
	Id        string `gorm:"primaryKey"`
	Email     string `gorm:"type:varchar(128);not null"`
	Password  string `gorm:"type:text;not null"`
	CreatedAt int64  `gorm:"type:int;not null"`
	UpdatedAt *int64 `gorm:"type:int;default:null"`
	DeletedAt *int64 `gorm:"type:int;default:null"`
}

func (Account) TableName() string {
	return "accounts"
}

func (m *Account) GetInfo(invoices []*invoice.InvoiceInfo) *payload.AccountInfo {
	invoiceInfos := make([]*payload.InvoiceInfo, 0)

	for _, invoice := range invoices {
		invoiceInfos = append(invoiceInfos, &payload.InvoiceInfo{
			Id:        invoice.GetId(),
			Email:     invoice.GetEmail(),
			Amount:    invoice.GetAmount(),
			CreatedAt: invoice.GetCreatedAt(),
		})
	}

	return &payload.AccountInfo{
		Id:        m.Id,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
		Invoices:  invoiceInfos,
	}
}
