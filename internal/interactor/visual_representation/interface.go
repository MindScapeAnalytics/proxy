package visualrepresentation

import "context"

type VisualRepresentationRepository interface {
	GetTestTemplateBySlug(ctx context.Context, slug string) ([]byte, error)
	GetTestingResultByAccountID(ctx context.Context, accountId string) ([]byte, error)
}
