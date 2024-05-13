package psychologytesting

import (
	"context"
	"encoding/json"

	api_entity "github.com/MindScapeAnalytics/grpc-api/core/client/entity"
	"github.com/MindScapeAnalytics/proxy/internal/entity"
	visualRepresentation "github.com/MindScapeAnalytics/proxy/internal/interactor/visual_representation"
)

type CoreInteractor struct {
	visualRepresentation visualRepresentation.VisualRepresentationRepository
	coreRepository       CoreRepository
}

type CoreInteractorOpts struct {
	VisualRepresentation visualRepresentation.VisualRepresentationRepository
	CoreRepository       CoreRepository
}

func NewCoreInteractor(ctx context.Context, opts CoreInteractorOpts) (*CoreInteractor, error) {
	return &CoreInteractor{
		visualRepresentation: opts.VisualRepresentation,
		coreRepository:       opts.CoreRepository,
	}, nil
}

func (interactor CoreInteractor) AddUser(ctx context.Context, user api_entity.User) error {
	if err := interactor.coreRepository.AddUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) AddEvent(ctx context.Context, event api_entity.Event, user api_entity.User) error {
	if err := interactor.coreRepository.AddEvent(ctx, event, user); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) AddFriend(ctx context.Context, user, userFriend api_entity.User) error {
	if err := interactor.coreRepository.AddFriend(ctx, user, userFriend); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) RemoveFriend(ctx context.Context, user, userFriend api_entity.User) error {
	if err := interactor.coreRepository.RemoveFriend(ctx, user, userFriend); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) AddEventToUser(ctx context.Context, user api_entity.User, event api_entity.Event) error {
	if err := interactor.coreRepository.AddEventToUser(ctx, user, event); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) AddCognitiveSpecificationToUser(ctx context.Context, user api_entity.User, cs api_entity.CognitiveSpecification) error {
	var (
		test entity.TestRawResult
	)

	res, err := interactor.visualRepresentation.GetTestingResultByAccountID(ctx, user.Id)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(res, &test); err != nil {
		return err
	}

	arr := make([]api_entity.Accentuation, 0)
	for k, v := range test.Raw {
		if v.Accentuated {
			arr = append(arr, api_entity.Accentuation{
				Type: k,
			})
		}
	}

	cs.Accentuations = arr

	if err := interactor.coreRepository.AddCognitiveSpecificationToUser(ctx, user, cs); err != nil {
		return err
	}

	return nil
}

func (interactor CoreInteractor) UpdateCognitiveSpecificationUser(ctx context.Context, user api_entity.User, cs api_entity.CognitiveSpecification) error {
	if err := interactor.coreRepository.UpdateCognitiveSpecificationUser(ctx, user, cs); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) UpdateAdditionalUserEventInfo(ctx context.Context, user api_entity.User, event api_entity.Event) error {
	if err := interactor.coreRepository.UpdateAdditionalUserEventInfo(ctx, user, event); err != nil {
		return err
	}
	return nil
}

func (interactor CoreInteractor) GetUserEventList(ctx context.Context, user api_entity.User, limit int) ([]api_entity.Event, error) {
	res, err := interactor.coreRepository.GetUserEventList(ctx, user, limit)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (interactor CoreInteractor) GetEventInfo(ctx context.Context, eventId string) (api_entity.Event, error) {
	res, err := interactor.coreRepository.GetEventInfo(ctx, eventId)
	if err != nil {
		return api_entity.Event{}, err
	}
	return res, nil
}

func (interactor CoreInteractor) AddEventInfo(ctx context.Context, event api_entity.Event, user api_entity.User) error {
	err := interactor.coreRepository.AddEventInfo(ctx, event, user)
	if err != nil {
		return err
	}
	return nil
}
