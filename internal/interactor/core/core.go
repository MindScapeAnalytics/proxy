package psychologytesting

import (
	"context"

	"github.com/MindScapeAnalytics/grpc-api/core/client/entity"
)

type CoreInteractor struct {
	coreRepository CoreRepository
}

type CoreInteractorOpts struct {
	CoreRepository CoreRepository
}

func NewCoreInteractor(ctx context.Context, opts CoreInteractorOpts) (*CoreInteractor, error) {
	return &CoreInteractor{
		coreRepository: opts.CoreRepository,
	}, nil
}

func (interactor CoreInteractor) AddUser(ctx context.Context, user entity.User) error {
	if err := interactor.coreRepository.AddUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) AddEvent(ctx context.Context, event entity.Event, user entity.User) error {
	if err := interactor.coreRepository.AddEvent(ctx, event, user); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) AddFriend(ctx context.Context, user, userFriend entity.User) error {
	if err := interactor.coreRepository.AddFriend(ctx, user, userFriend); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) RemoveFriend(ctx context.Context, user, userFriend entity.User) error {
	if err := interactor.coreRepository.RemoveFriend(ctx, user, userFriend); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) AddEventToUser(ctx context.Context, user entity.User, event entity.Event) error {
	if err := interactor.coreRepository.AddEventToUser(ctx, user, event); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) AddCognitiveSpecificationToUser(ctx context.Context, user entity.User, cs entity.CognitiveSpecification) error {
	if err := interactor.coreRepository.AddCognitiveSpecificationToUser(ctx, user, cs); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) UpdateCognitiveSpecificationUser(ctx context.Context, user entity.User, cs entity.CognitiveSpecification) error {
	if err := interactor.coreRepository.UpdateCognitiveSpecificationUser(ctx, user, cs); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) UpdateAdditionalUserEventInfo(ctx context.Context, user entity.User, event entity.Event) error {
	if err := interactor.coreRepository.UpdateAdditionalUserEventInfo(ctx, user, event); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) GetUserEventList(ctx context.Context, user entity.User) ([][]byte, error) {
	res, err := interactor.coreRepository.GetUserEventList(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
