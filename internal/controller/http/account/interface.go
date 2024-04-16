package user

import (
	"context"

	"github.com/MindScapeAnalytics/proxy/internal/entity"
)

type AccountInteractor interface {
	Login(ctx context.Context, account entity.Account) (string, error)
	Registry(ctx context.Context, account entity.Account) (bool, error)
}
