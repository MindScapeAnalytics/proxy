package psychologytesting

import "context"

type PsychologyTestingInteractor struct {
	psychologyTestingRepository PsychologyTestingRepository
}

type PsychologyTestingInteractorOpts struct {
	PsychologyTestingRepository PsychologyTestingRepository
}

func NewPsychologyTestingInteractor(ctx context.Context, opts PsychologyTestingInteractorOpts) (*PsychologyTestingInteractor, error) {
	return &PsychologyTestingInteractor{
		psychologyTestingRepository: opts.PsychologyTestingRepository,
	}, nil
}

func (interactor PsychologyTestingInteractor) SendTestingData(ctx context.Context, answers []bool, accountId string) error {
	if err := interactor.psychologyTestingRepository.SendTestingData(ctx, answers, accountId); err != nil {
		return err
	}
	return nil
}
