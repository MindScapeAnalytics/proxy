package psychologytesting

import "context"

type PsychologyTestingRepository interface {
	SendTestingData(ctx context.Context, answers []bool) error
}
