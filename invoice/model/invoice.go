package model

import (
	"github.com/syauqeesy/accounting-service/invoice/payload"
	"github.com/syauqeesy/accounting-service/proto/compiled/account"
)

type Invoice struct {
	Id        string  `gorm:"primaryKey"`
	UserId    string  `gorm:"type:char(32);not null"`
	Email     string  `gorm:"type:varchar(128);not null"`
	Amount    float32 `gorm:"type:double;not null"`
	CreatedAt int64   `gorm:"type:int;not null"`
	UpdatedAt *int64  `gorm:"type:int;default:null"`
	DeletedAt *int64  `gorm:"type:int;default:null"`
}

func (Invoice) TableName() string {
	return "invoices"
}

func (m *Invoice) GetInfo(account *account.AccountInfo) *payload.InvoiceInfo {
	return &payload.InvoiceInfo{
		Id: m.Id,
		Account: &payload.AccountInfo{
			Id:        account.GetId(),
			Email:     account.GetEmail(),
			CreatedAt: account.GetCreatedAt(),
		},
		Email:     m.Email,
		Amount:    m.Amount,
		CreatedAt: m.CreatedAt,
	}
}
