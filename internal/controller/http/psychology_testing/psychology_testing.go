package psychologytesting

import (
	"context"
	"time"

	"github.com/MindScapeAnalytics/proxy/pkg/logger"
	"github.com/MindScapeAnalytics/proxy/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type PsychologyTestingController struct {
	psychologyTestingInteractor PsychologyTestingInteractor
	logger                      logger.LoggerUC
}

type PsychologyTestingCtrlOpts struct {
	PsychologyTestingInteractor PsychologyTestingInteractor
	Logger                      logger.LoggerUC
}

func NewAccountController(ctx context.Context, opts PsychologyTestingCtrlOpts) (PsychologyTestingController, error) {
	return PsychologyTestingController{
		psychologyTestingInteractor: opts.PsychologyTestingInteractor,
		logger:                      opts.Logger,
	}, nil
}

func (controller PsychologyTestingController) SendTestingData() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		var (
			answers struct {
				Answers []bool `json:"answers"`
			}
		)

		err := utils.ReadRequest(ctx, &answers)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		if err := controller.psychologyTestingInteractor.SendTestingData(ctx.Context(), answers.Answers); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return ctx.SendStatus(fiber.StatusAccepted)
	}
}
