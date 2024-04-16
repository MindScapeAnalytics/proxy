package account

import (
	"context"
	"errors"

	api "github.com/MindScapeAnalytics/grpc-api/authentication/client"
	accountRepo "github.com/MindScapeAnalytics/proxy/internal/interactor/account"
)

const (
	authenticationService = "authenticationService"
)

type AccountRepOpts struct {
	Type                  string
	AuthenticationService api.AuthenticationService
}

func NewAccountRepository(ctx context.Context, opts AccountRepOpts) (accountRepo.AccountRepository, error) {
	switch opts.Type {
	case authenticationService:
		return newAuthenticationRepository(ctx, opts)
	default:
		return nil, errors.New("неподдерживаемый тип репозитория")
	}
}
