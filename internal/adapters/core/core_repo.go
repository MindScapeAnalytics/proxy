package core

import (
	"context"
	"errors"

	api "github.com/MindScapeAnalytics/grpc-api/core/client"
	interactor "github.com/MindScapeAnalytics/proxy/internal/interactor/core"
)

const (
	coreService = "coreService"
)

type CoreRepositoryOpts struct {
	Type        string
	CoreService api.CoreService
}

func NewCoreRepository(ctx context.Context, opts CoreRepositoryOpts) (interactor.CoreRepository, error) {
	switch opts.Type {
	case coreService:
		return newCoreRepositoryService(ctx, opts)
	default:
		return nil, errors.New("неподдерживаемый тип репозитория")
	}
}
