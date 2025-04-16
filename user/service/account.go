package service

import "github.com/syauqeesy/accounting-service/user/payload"

type AccountService interface {
	List() ([]*payload.AccountInfo, error)
}

type accountService service

func (s *accountService) List() ([]*payload.AccountInfo, error) {
	accountInfos := make([]*payload.AccountInfo, 0)

	accounts, err := s.Repository.Account.Select()
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		accountInfos = append(accountInfos, account.GetInfo())
	}

	return accountInfos, nil
}
