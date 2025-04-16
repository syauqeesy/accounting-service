package model

import "github.com/syauqeesy/accounting-service/user/payload"

type Account struct {
	Id        string `gorm:"primaryKey"`
	Email     string `gorm:"type:varchar(128);not null"`
	Password  string `gorm:"type:text;not null"`
	CreatedAt int    `gorm:"type:int;not null"`
	UpdatedAt *int   `gorm:"type:int;default:null"`
	DeletedAt *int   `gorm:"type:int;default:null"`
}

func (Account) TableName() string {
	return "accounts"
}

func (a *Account) GetInfo() *payload.AccountInfo {
	return &payload.AccountInfo{
		Id:        a.Id,
		Email:     a.Email,
		CreatedAt: a.CreatedAt,
	}
}
