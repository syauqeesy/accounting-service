package service

import (
	"context"

	"github.com/syauqeesy/accounting-service/proto/compiled/invoice"
	"github.com/syauqeesy/accounting-service/user/payload"
)

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
		result, err := s.GRPCOutbound.Invoice.SelectByUserId(context.Background(), &invoice.SelectByUserIdRequest{
			UserId: account.Id,
		})
		if err != nil {
			return nil, err
		}

		accountInfos = append(accountInfos, account.GetInfo(result.GetInvoices()))
	}

	return accountInfos, nil
}
