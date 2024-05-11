package visualrepresentation

import (
	"context"
)

type VisualRepresentationInteractor interface {
	GetTestingResultByAccountID(ctx context.Context, accountId string) ([]byte, error)
	GetTestTemplateBySlug(ctx context.Context, slug string) ([]byte, error)
}
