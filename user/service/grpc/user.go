package grpc_service

import (
	"context"

	"github.com/syauqeesy/accounting-service/proto/compiled/account"
	"github.com/syauqeesy/accounting-service/user/model"
)

type accountService struct {
	grpcService
	account.UnimplementedAccountServiceServer
}

func (s *accountService) getAccountInfo(accountModel *model.Account) *account.AccountInfo {
	return &account.AccountInfo{
		Id:        accountModel.Id,
		Email:     accountModel.Email,
		CreatedAt: accountModel.CreatedAt,
	}
}

func (s *accountService) SelectById(ctx context.Context, request *account.SelectByIdRequest) (*account.AccountInfo, error) {
	account, err := s.Repository.Account.SelectById(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return s.getAccountInfo(account), nil
}
