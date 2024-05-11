package psychologytesting

import (
	"context"
)

type PsychologyTestingInteractor interface {
	SendTestingData(ctx context.Context, answers []bool) error
}
