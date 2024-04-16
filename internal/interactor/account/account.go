package account

import (
	"context"

	"github.com/MindScapeAnalytics/proxy/internal/entity"
)

type AccountInteractor struct {
	accountRepository AccountRepository
}

type AccountIntrOpts struct {
	AccountRepository AccountRepository
}

func NewAccountInteractor(ctx context.Context, opts AccountIntrOpts) (*AccountInteractor, error) {
	return &AccountInteractor{
		accountRepository: opts.AccountRepository,
	}, nil
}

func (interactor AccountInteractor) Login(ctx context.Context, account entity.Account) (string, error) {
	token, err := interactor.accountRepository.GetToken(ctx, account.Login, account.Password)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (interactor AccountInteractor) Registry(ctx context.Context, account entity.Account) (bool, error) {
	ok, err := interactor.accountRepository.Registry(ctx, account.Login, account.Email, account.Password)
	if err != nil {
		return false, err
	}
	return ok, nil
}
