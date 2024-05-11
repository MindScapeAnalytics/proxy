package core

import (
	"context"

	api "github.com/MindScapeAnalytics/grpc-api/psychology_testing/client"
)

type psychologyTestingRepository struct {
	psychologyTestingService api.PsychologyTestingService
}

func newPsychologyTestingService(ctx context.Context, opts PsychologyTestingRepositoryOpts) (*psychologyTestingRepository, error) {
	return &psychologyTestingRepository{
		psychologyTestingService: opts.PsychologyTestingService,
	}, nil
}

func (repository psychologyTestingRepository) SendTestingData(ctx context.Context, answers []bool) error {
	if err := repository.psychologyTestingService.SendTestingData(ctx, answers); err != nil {
		return err
	}
	return nil
}
