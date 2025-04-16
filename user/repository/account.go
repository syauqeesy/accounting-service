package repository

import "github.com/syauqeesy/accounting-service/user/model"

type AccountRepository interface {
	Select() ([]*model.Account, error)
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
