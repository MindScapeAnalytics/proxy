package user

import (
	"context"
	"time"

	"github.com/MindScapeAnalytics/proxy/internal/entity"
	"github.com/MindScapeAnalytics/proxy/pkg/logger"
	"github.com/MindScapeAnalytics/proxy/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	accountInteractor AccountInteractor
	logger            logger.LoggerUC
}

type AccountCtrlOpts struct {
	AccountInteractor AccountInteractor
	Logger            logger.LoggerUC
}

func NewAccountController(ctx context.Context, opts AccountCtrlOpts) (AccountController, error) {
	return AccountController{
		accountInteractor: opts.AccountInteractor,
		logger:            opts.Logger,
	}, nil
}

func (controller AccountController) Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		var (
			account entity.Account
		)

		err := utils.ReadRequest(ctx, &account)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		token, err := controller.accountInteractor.Login(ctx.Context(), account)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusAccepted).JSON(token)
	}
}

func (controller AccountController) Registry() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		var (
			account entity.Account
		)

		err := utils.ReadRequest(ctx, &account)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		ok, err := controller.accountInteractor.Registry(ctx.Context(), account)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		if !ok {
			return ctx.Status(fiber.StatusBadRequest).JSON("something whet wrong")
		}
		return ctx.SendStatus(fiber.StatusAccepted)
	}
}
