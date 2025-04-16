package repository

import "gorm.io/gorm"

type repository struct {
	Database *gorm.DB
}

type Repository struct {
	Invoice InvoiceRepository
}

func New(database *gorm.DB) *Repository {
	repository := &repository{
		Database: database,
	}

	return &Repository{
		Invoice: (*invoiceRepository)(repository),
	}
}
