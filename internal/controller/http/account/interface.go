package user

import (
	"context"

	"github.com/MindScapeAnalytics/proxy/internal/entity"
)

type AccountInteractor interface {
	Login(ctx context.Context, account entity.Account) ([]byte, error)
	Registry(ctx context.Context, account entity.Account) (entity.Account, error)
}
