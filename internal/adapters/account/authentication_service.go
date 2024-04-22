package account

import (
	"context"
	"fmt"

	api "github.com/MindScapeAnalytics/grpc-api/authentication/client"
)

type authenticationServiceRepository struct {
	authenticationService api.AuthenticationService
}

func newAuthenticationRepository(ctx context.Context, opts AccountRepOpts) (*authenticationServiceRepository, error) {

	return &authenticationServiceRepository{
		authenticationService: opts.AuthenticationService,
	}, nil
}

func (repository authenticationServiceRepository) GetToken(ctx context.Context, login, password string) ([]byte, error) {
	token, err := repository.authenticationService.GetToken(ctx, login, password)
	if err != nil {
		return nil, fmt.Errorf("authenticationServiceRepository.GetToken(): error: %s", err.Error())
	}
	return token, nil
}

func (repository authenticationServiceRepository) Registry(ctx context.Context, login, email, password string) (bool, error) {
	_, err := repository.authenticationService.Register(ctx, login, password, email)
	if err != nil {
		return false, fmt.Errorf("authenticationServiceRepository.GetToken(): error: %s", err.Error())
	}
	return true, nil
}
