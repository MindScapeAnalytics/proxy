package core

import (
	"context"

	api "github.com/MindScapeAnalytics/grpc-api/core/client"
	"github.com/MindScapeAnalytics/grpc-api/core/client/entity"
)

type coreRepository struct {
	coreService api.CoreService
}

func newCoreRepositoryService(ctx context.Context, opts CoreRepositoryOpts) (*coreRepository, error) {
	return &coreRepository{
		coreService: opts.CoreService,
	}, nil
}

func (repo coreRepository) AddUser(ctx context.Context, user entity.User) error {
	err := repo.coreService.AddUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (repo coreRepository) AddEvent(ctx context.Context, event entity.Event, user entity.User) error {
	err := repo.coreService.AddEvent(ctx, event, user)
	if err != nil {
		return err
	}
	return nil
}

func (repo coreRepository) AddFriend(ctx context.Context, user, userFriend entity.User) error {
	err := repo.coreService.AddFriend(ctx, user, userFriend)
	if err != nil {
		return err
	}
	return nil
}

func (repo coreRepository) RemoveFriend(ctx context.Context, user, userFriend entity.User) error {
	err := repo.coreService.RemoveFriend(ctx, user, userFriend)
	if err != nil {
		return err
	}
	return nil
}

func (repo coreRepository) AddEventToUser(ctx context.Context, user entity.User, event entity.Event) error {
	err := repo.coreService.AddEventToUser(ctx, user, event)
	if err != nil {
		return err
	}
	return nil
}

func (repo coreRepository) AddCognitiveSpecificationToUser(ctx context.Context, user entity.User, cs entity.CognitiveSpecification) error {
	err := repo.coreService.AddCognitiveSpecificationToUser(ctx, user, cs)
	if err != nil {
		return err
	}
	return nil
}

func (repo coreRepository) UpdateCognitiveSpecificationUser(ctx context.Context, user entity.User, cs entity.CognitiveSpecification) error {
	err := repo.coreService.UpdateCognitiveSpecificationUser(ctx, user, cs)
	if err != nil {
		return err
	}
	return nil
}

func (repo coreRepository) UpdateAdditionalUserEventInfo(ctx context.Context, user entity.User, event entity.Event) error {
	err := repo.coreService.UpdateAdditionalUserEventInfo(ctx, user, event)
	if err != nil {
		return err
	}
	return nil
}

func (repo coreRepository) GetUserEventList(ctx context.Context, user entity.User, limit int) ([]entity.Event, error) {
	res, err := repo.coreService.GetUserEventList(ctx, user, limit)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repo coreRepository) GetEventInfo(ctx context.Context, eventId string) (entity.Event, error) {
	res, err := repo.coreService.GetEventInfo(ctx, entity.Event{
		Id: eventId,
	})
	if err != nil {
		return entity.Event{}, err
	}
	return res, nil
}

func (repo coreRepository) AddEventInfo(ctx context.Context, event entity.Event, user entity.User) error {
	err := repo.coreService.AddEventInfo(ctx, event, user)
	if err != nil {
		return err
	}
	return nil
}
