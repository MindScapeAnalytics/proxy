package visualrepresentation

import (
	"context"
	"errors"

	api "github.com/MindScapeAnalytics/grpc-api/visual_representation/client"
	visualRepresentation "github.com/MindScapeAnalytics/proxy/internal/interactor/visual_representation"
)

const (
	visualRepresentationService = "visualRepresentationService"
)

type VisualRepresentationRepOpts struct {
	Type                        string
	VisualRepresentationService api.VisualRepresentationService
}

func NewVisualRepresentation(ctx context.Context, opts VisualRepresentationRepOpts) (visualRepresentation.VisualRepresentationRepository, error) {
	switch opts.Type {
	case visualRepresentationService:
		return newVisualRepresentationRepository(ctx, opts)
	default:
		return nil, errors.New("неподдерживаемый тип репозитория")
	}
}
