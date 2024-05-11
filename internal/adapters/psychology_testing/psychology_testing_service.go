package psychologytesting

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

func (repository psychologyTestingRepository) SendTestingData(ctx context.Context, answers []bool, accountId string) error {
	if err := repository.psychologyTestingService.SendTestingData(ctx, answers, accountId); err != nil {
		return err
	}
	return nil
}
