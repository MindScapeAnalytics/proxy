package visualrepresentation

import (
	"context"

	api "github.com/MindScapeAnalytics/grpc-api/visual_representation/client"
)

type visualRepresentationRepository struct {
	visualRepresentationService api.VisualRepresentationService
}

func newVisualRepresentationRepository(ctx context.Context, opts VisualRepresentationRepOpts) (*visualRepresentationRepository, error) {
	return &visualRepresentationRepository{
		visualRepresentationService: opts.VisualRepresentationService,
	}, nil
}

func (repository visualRepresentationRepository) GetTestingResultByAccountID(ctx context.Context, accountId string) ([]byte, error) {
	res, err := repository.visualRepresentationService.GetTestingResultByAccountID(ctx, accountId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repository visualRepresentationRepository) GetTestTemplateBySlug(ctx context.Context, slug string) ([]byte, error) {
	res, err := repository.visualRepresentationService.GetTestTemplateBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	return res, nil
}
