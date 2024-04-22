package account

import (
	"context"
)

type AccountRepository interface {
	GetToken(ctx context.Context, login, password string) ([]byte, error)
	Registry(ctx context.Context, login, email, password string) (bool, error)
}
