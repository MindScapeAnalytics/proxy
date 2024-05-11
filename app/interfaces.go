package app

import (
	"context"

	authenticationService "github.com/MindScapeAnalytics/grpc-api/authentication/api/v1"
)

// UserHTTPController ...
type AccountController interface {
	GetToken(ctx context.Context, req *authenticationService.GetTokenRequest) (*authenticationService.GetTokenResponse, error)
}
