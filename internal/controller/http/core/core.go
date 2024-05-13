package user

import (
	"context"
	"strconv"
	"time"

	api_entity "github.com/MindScapeAnalytics/grpc-api/core/client/entity"
	"github.com/MindScapeAnalytics/proxy/pkg/logger"
	"github.com/MindScapeAnalytics/proxy/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type CoreController struct {
	coreInteractor CoreInteractor
	logger         logger.LoggerUC
}

type CoreControllerOpts struct {
	CoreInteractor CoreInteractor
	Logger         logger.LoggerUC
}

func NewCoreController(ctx context.Context, opts CoreControllerOpts) (CoreController, error) {
	return CoreController{
		coreInteractor: opts.CoreInteractor,
		logger:         opts.Logger,
	}, nil
}

func (controller CoreController) AddUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		var (
			user struct {
				EventActionType []string
				Accentuations   []string
			}
		)

		err := utils.ReadRequest(ctx, &user)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		arr := make([]api_entity.Accentuation, 0)
		for _, v := range user.Accentuations {
			arr = append(arr, api_entity.Accentuation{
				Type: v,
			})
		}
		err = controller.coreInteractor.AddUser(ctx.Context(), api_entity.User{
			Id: ctx.Locals("accountId").(string),
			EventActions: api_entity.EventActions{
				Type: user.EventActionType,
			},
			CognitiveSpecification: api_entity.CognitiveSpecification{
				Accentuations: arr,
			},
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusAccepted)
	}
}

func (controller CoreController) AddEvent() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		var (
			event struct {
				Name        string `json:"name"`
				Description string `json:"description"`
				Data        string `json:"data"`
			}
		)

		err := utils.ReadRequest(ctx, &event)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		err = controller.coreInteractor.AddEvent(
			ctx.Context(),
			api_entity.Event{
				Name:        event.Name,
				Description: event.Description,
				Data:        []byte(event.Data),
				UserId:      ctx.Locals("accountId").(string),
			}, api_entity.User{
				Id: ctx.Locals("accountId").(string),
			})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusAccepted)
	}
}

func (controller CoreController) AddFriend() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		var (
			user struct {
				EventActionType []string
				Accentuations   []string
			}
			userFriend struct {
				EventActionType []string
				Accentuations   []string
			}
		)

		err := utils.ReadRequest(ctx, &user)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		err = utils.ReadRequest(ctx, &userFriend)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		arr := make([]api_entity.Accentuation, 0)
		for _, v := range user.Accentuations {
			arr = append(arr, api_entity.Accentuation{
				Type: v,
			})
		}
		arrFriend := make([]api_entity.Accentuation, 0)
		for _, v := range user.Accentuations {
			arrFriend = append(arrFriend, api_entity.Accentuation{
				Type: v,
			})
		}

		err = controller.coreInteractor.AddFriend(
			ctx.Context(),
			api_entity.User{
				Id: ctx.Locals("accountId").(string),
				EventActions: api_entity.EventActions{
					Type: user.EventActionType,
				},
				CognitiveSpecification: api_entity.CognitiveSpecification{
					Accentuations: arr,
				},
			},
			api_entity.User{
				Id: ctx.Locals("accountId").(string),
				EventActions: api_entity.EventActions{
					Type: userFriend.EventActionType,
				},
				CognitiveSpecification: api_entity.CognitiveSpecification{
					Accentuations: arrFriend,
				},
			},
		)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusAccepted)
	}
}

func (controller CoreController) RemoveFriend() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		var (
			user struct {
				EventActionType []string
				Accentuations   []string
			}
			userFriend struct {
				EventActionType []string
				Accentuations   []string
			}
		)

		err := utils.ReadRequest(ctx, &user)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		err = utils.ReadRequest(ctx, &userFriend)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		arr := make([]api_entity.Accentuation, 0)
		for _, v := range user.Accentuations {
			arr = append(arr, api_entity.Accentuation{
				Type: v,
			})
		}
		arrFriend := make([]api_entity.Accentuation, 0)
		for _, v := range user.Accentuations {
			arrFriend = append(arrFriend, api_entity.Accentuation{
				Type: v,
			})
		}

		err = controller.coreInteractor.RemoveFriend(
			ctx.Context(),
			api_entity.User{
				Id: ctx.Locals("accountId").(string),
				EventActions: api_entity.EventActions{
					Type: user.EventActionType,
				},
				CognitiveSpecification: api_entity.CognitiveSpecification{
					Accentuations: arr,
				},
			},
			api_entity.User{
				Id: ctx.Locals("accountId").(string),
				EventActions: api_entity.EventActions{
					Type: userFriend.EventActionType,
				},
				CognitiveSpecification: api_entity.CognitiveSpecification{
					Accentuations: arrFriend,
				},
			},
		)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusAccepted)
	}
}

func (controller CoreController) AddEventToUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		var (
			event struct {
				Id string `json:"id"`
			}
			user struct {
				EventActionType []string
				Accentuations   []string
			}
		)

		err := utils.ReadRequest(ctx, &event)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		err = utils.ReadRequest(ctx, &user)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		arr := make([]api_entity.Accentuation, 0)
		for _, v := range user.Accentuations {
			arr = append(arr, api_entity.Accentuation{
				Type: v,
			})
		}
		err = controller.coreInteractor.AddEventToUser(
			ctx.Context(),
			api_entity.User{
				Id: ctx.Locals("accountId").(string),
				EventActions: api_entity.EventActions{
					Type: user.EventActionType,
				},
				CognitiveSpecification: api_entity.CognitiveSpecification{
					Accentuations: arr,
				},
			},
			api_entity.Event{
				Id: event.Id,
			})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusAccepted)
	}
}

func (controller CoreController) AddCognitiveSpecificationToUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		err := controller.coreInteractor.AddCognitiveSpecificationToUser(
			ctx.Context(),
			api_entity.User{
				Id: ctx.Locals("accountId").(string),
			},
			api_entity.CognitiveSpecification{},
		)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return ctx.SendStatus(fiber.StatusAccepted)
	}
}

func (controller CoreController) UpdateCognitiveSpecificationUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		var (
			user struct {
				EventActionType []string
				Accentuations   []string
			}
			cs struct {
				Accentuations []string
			}
		)

		err := utils.ReadRequest(ctx, &user)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		arr := make([]api_entity.Accentuation, 0)
		for _, v := range cs.Accentuations {
			arr = append(arr, api_entity.Accentuation{
				Type: v,
			})
		}

		err = controller.coreInteractor.UpdateCognitiveSpecificationUser(
			ctx.Context(),
			api_entity.User{
				Id: ctx.Locals("accountId").(string),
				EventActions: api_entity.EventActions{
					Type: user.EventActionType,
				},
			},
			api_entity.CognitiveSpecification{
				Accentuations: arr,
			},
		)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return ctx.SendStatus(fiber.StatusAccepted)
	}
}

func (controller CoreController) UpdateAdditionalUserEventInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		var (
			event struct {
				Id string `json:"id"`
			}
			user struct {
				EventActionType []string
				Accentuations   []string
			}
		)

		err := utils.ReadRequest(ctx, &event)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		err = utils.ReadRequest(ctx, &user)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		arr := make([]api_entity.Accentuation, 0)
		for _, v := range user.Accentuations {
			arr = append(arr, api_entity.Accentuation{
				Type: v,
			})
		}
		err = controller.coreInteractor.UpdateAdditionalUserEventInfo(
			ctx.Context(),
			api_entity.User{
				Id: ctx.Locals("accountId").(string),
				EventActions: api_entity.EventActions{
					Type: user.EventActionType,
				},
				CognitiveSpecification: api_entity.CognitiveSpecification{
					Accentuations: arr,
				},
			},
			api_entity.Event{
				Id: event.Id,
			},
		)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusAccepted)
	}
}

func (controller CoreController) GetUserEventList() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())
		limit := ctx.Query("limit", "10")
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Param must be an integer")
		}
		res, err := controller.coreInteractor.GetUserEventList(ctx.Context(), api_entity.User{
			Id: ctx.Locals("accountId").(string),
		}, limitInt)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		// arr := make([]entity.Event, 0)
		// for _, v := range res {
		// 	arr = append(arr, entity.Event{
		// 		v.
		// 	})
		// }
		return ctx.Status(fiber.StatusAccepted).JSON(res)
	}
}

func (controller CoreController) GetEventInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		id := ctx.Params("id")

		res, err := controller.coreInteractor.GetEventInfo(ctx.Context(), id)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusAccepted).JSON(res)
	}
}

func (controller CoreController) AddEventInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		var (
			event struct {
				Id          string `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Data        string `json:"data"`
				UserId      string `json:"userId"`
			}
		)

		err := utils.ReadRequest(ctx, &event)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		userId := ctx.Locals("accountId").(string)

		err = controller.coreInteractor.AddEventInfo(
			ctx.Context(),
			api_entity.Event{
				Id:          event.Id,
				Name:        event.Name,
				Description: event.Description,
				Data:        []byte(event.Data),
				UserId:      event.UserId,
			},
			api_entity.User{
				Id: userId,
			},
		)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusAccepted)
	}
}
