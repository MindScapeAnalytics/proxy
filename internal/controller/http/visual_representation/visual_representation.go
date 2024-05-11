package visualrepresentation

import (
	"context"
	"time"

	"github.com/MindScapeAnalytics/proxy/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type VisualRepresentationController struct {
	visualRepresentationInteractor VisualRepresentationInteractor
	logger                         logger.LoggerUC
}

type VisualRepresentationControllerOpts struct {
	VisualRepresentationInteractor VisualRepresentationInteractor
	Logger                         logger.LoggerUC
}

func NewVisualRepresentationController(ctx context.Context, opts VisualRepresentationControllerOpts) (VisualRepresentationController, error) {
	return VisualRepresentationController{
		visualRepresentationInteractor: opts.VisualRepresentationInteractor,
		logger:                         opts.Logger,
	}, nil
}

func (controller VisualRepresentationController) GetTestingResultByAccountID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		id := ctx.Params("id")

		res, err := controller.visualRepresentationInteractor.GetTestingResultByAccountID(ctx.Context(), id)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusAccepted).JSON(res)
	}
}

func (controller VisualRepresentationController) GetTestTemplateBySlug() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())

		slug := ctx.Params("slug")

		res, err := controller.visualRepresentationInteractor.GetTestTemplateBySlug(ctx.Context(), slug)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusAccepted).JSON(res)
	}
}
