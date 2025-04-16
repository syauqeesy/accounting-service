package model

import "github.com/syauqeesy/accounting-service/invoice/payload"

type Invoice struct {
	Id        string  `gorm:"primaryKey"`
	UserId    string  `gorm:"type:char(32);not null"`
	Email     string  `gorm:"type:varchar(128);not null"`
	Amount    float64 `gorm:"type:double;not null"`
	CreatedAt int     `gorm:"type:int;not null"`
	UpdatedAt *int    `gorm:"type:int;default:null"`
	DeletedAt *int    `gorm:"type:int;default:null"`
}

func (Invoice) TableName() string {
	return "invoices"
}

func (m *Invoice) GetInfo() *payload.InvoiceInfo {
	return &payload.InvoiceInfo{
		Id:        m.Id,
		Email:     m.Email,
		Amount:    m.Amount,
		CreatedAt: m.CreatedAt,
	}
}
