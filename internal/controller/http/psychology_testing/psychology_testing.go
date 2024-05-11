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

func NewPsychologyTestingController(ctx context.Context, opts PsychologyTestingCtrlOpts) (PsychologyTestingController, error) {
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
				Answers   []bool `json:"answers"`
				AccountId string `json:"accountId"`
			}
		)

		err := utils.ReadRequest(ctx, &answers)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		answers.AccountId = ctx.Locals("accountId").(string)

		if err := controller.psychologyTestingInteractor.SendTestingData(ctx.Context(), answers.Answers, answers.AccountId); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return ctx.SendStatus(fiber.StatusAccepted)
	}
}
