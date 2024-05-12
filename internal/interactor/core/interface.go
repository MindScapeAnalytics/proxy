package psychologytesting

import (
	"context"

	"github.com/MindScapeAnalytics/grpc-api/core/client/entity"
)

type CoreRepository interface {
	AddUser(ctx context.Context, user entity.User) error
	AddEvent(ctx context.Context, event entity.Event, user entity.User) error
	AddFriend(ctx context.Context, user, userFriend entity.User) error
	RemoveFriend(ctx context.Context, user, userFriend entity.User) error
	AddEventToUser(ctx context.Context, user entity.User, event entity.Event) error
	AddCognitiveSpecificationToUser(ctx context.Context, user entity.User, cs entity.CognitiveSpecification) error
	UpdateCognitiveSpecificationUser(ctx context.Context, user entity.User, cs entity.CognitiveSpecification) error
	UpdateAdditionalUserEventInfo(ctx context.Context, user entity.User, event entity.Event) error
	GetUserEventList(ctx context.Context, user entity.User) ([][]byte, error)
}
