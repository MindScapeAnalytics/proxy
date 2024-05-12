package account

import (
	"context"

	"github.com/MindScapeAnalytics/proxy/internal/entity"
)

type AccountRepository interface {
	GetToken(ctx context.Context, login, password string) ([]byte, error)
	Registry(ctx context.Context, login, email, password string) (entity.Account, error)
}
