package visualrepresentation

import "context"

type VisualRepresentationInteractor struct {
	visualRepresentationRepository VisualRepresentationRepository
}

type VisualRepresentationInteractorOpts struct {
	VisualRepresentationRepository VisualRepresentationRepository
}

func NewVisualRepresentationInteractor(ctx context.Context, opts VisualRepresentationInteractorOpts) (*VisualRepresentationInteractor, error) {
	return &VisualRepresentationInteractor{
		visualRepresentationRepository: opts.VisualRepresentationRepository,
	}, nil
}

func (interactor VisualRepresentationInteractor) GetTestingResultByAccountID(ctx context.Context, accountId string) ([]byte, error) {
	res, err := interactor.visualRepresentationRepository.GetTestingResultByAccountID(ctx, accountId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (interactor VisualRepresentationInteractor) GetTestTemplateBySlug(ctx context.Context, slug string) ([]byte, error) {
	res, err := interactor.visualRepresentationRepository.GetTestTemplateBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	return res, nil
}
