package account

import (
	"context"

	api_entity "github.com/MindScapeAnalytics/grpc-api/core/client/entity"
	"github.com/MindScapeAnalytics/proxy/internal/entity"
	coreRepo "github.com/MindScapeAnalytics/proxy/internal/interactor/core"
)

type AccountInteractor struct {
	accountRepository AccountRepository
	coreRepository    coreRepo.CoreRepository
}

type AccountIntrOpts struct {
	AccountRepository AccountRepository
	CoreRepository    coreRepo.CoreRepository
}

func NewAccountInteractor(ctx context.Context, opts AccountIntrOpts) (*AccountInteractor, error) {
	return &AccountInteractor{
		coreRepository:    opts.CoreRepository,
		accountRepository: opts.AccountRepository,
	}, nil
}

func (interactor AccountInteractor) Login(ctx context.Context, account entity.Account) ([]byte, error) {
	token, err := interactor.accountRepository.GetToken(ctx, account.Login, account.Password)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (interactor AccountInteractor) Registry(ctx context.Context, account entity.Account) (entity.Account, error) {
	ok, err := interactor.accountRepository.Registry(ctx, account.Login, account.Email, account.Password)
	if err != nil {
		return entity.Account{}, err
	}

	if err := interactor.coreRepository.AddUser(ctx, api_entity.User{
		Id: ok.Id,
	}); err != nil {
		return entity.Account{}, err
	}
	return ok, nil
}
