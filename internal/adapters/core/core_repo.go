package core

import (
	"context"
	"errors"

	api "github.com/MindScapeAnalytics/grpc-api/psychology_testing/client"
	interactor "github.com/MindScapeAnalytics/proxy/internal/interactor/psychology_testing"
)

const (
	psychologyTestingService = "psychologyTestingService"
)

type PsychologyTestingRepositoryOpts struct {
	Type                     string
	PsychologyTestingService api.PsychologyTestingService
}

func NewPsychologyTestingRepository(ctx context.Context, opts PsychologyTestingRepositoryOpts) (interactor.PsychologyTestingRepository, error) {
	switch opts.Type {
	case psychologyTestingService:
		return newPsychologyTestingService(ctx, opts)
	default:
		return nil, errors.New("неподдерживаемый тип репозитория")
	}
}
