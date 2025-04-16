package repository

import (
	"context"

	"github.com/syauqeesy/accounting-service/user/model"
)

type AccountRepository interface {
	Select() ([]*model.Account, error)
	SelectById(ctx context.Context, id string) (*model.Account, error)
}

type accountRepository repository

func (r *accountRepository) Select() ([]*model.Account, error) {
	accounts := make([]*model.Account, 0)

	q := r.Database.Find(&accounts)
	if q.Error != nil {
		return nil, q.Error
	}

	return accounts, nil
}

func (r *accountRepository) SelectById(ctx context.Context, id string) (*model.Account, error) {
	account := &model.Account{}

	q := r.Database.WithContext(ctx).Where("id = ?", id).First(&account)
	if q.Error != nil {
		return nil, q.Error
	}

	return account, nil
}
