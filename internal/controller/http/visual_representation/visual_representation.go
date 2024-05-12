package visualrepresentation

import (
	"context"
	"encoding/json"
	"time"

	api_entity "github.com/MindScapeAnalytics/grpc-api/core/client/entity"
	coreIntr "github.com/MindScapeAnalytics/proxy/internal/controller/http/core"
	"github.com/MindScapeAnalytics/proxy/internal/entity"
	"github.com/MindScapeAnalytics/proxy/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type VisualRepresentationController struct {
	coreInteractor                 coreIntr.CoreInteractor
	visualRepresentationInteractor VisualRepresentationInteractor
	logger                         logger.LoggerUC
}

type VisualRepresentationControllerOpts struct {
	CoreInteractor                 coreIntr.CoreInteractor
	VisualRepresentationInteractor VisualRepresentationInteractor
	Logger                         logger.LoggerUC
}

func NewVisualRepresentationController(ctx context.Context, opts VisualRepresentationControllerOpts) (VisualRepresentationController, error) {
	return VisualRepresentationController{
		coreInteractor:                 opts.CoreInteractor,
		visualRepresentationInteractor: opts.VisualRepresentationInteractor,
		logger:                         opts.Logger,
	}, nil
}

func (controller VisualRepresentationController) GetTestingResultByAccountID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())
		var (
			test entity.TestRawResult
		)
		id := ctx.Locals("accountId").(string)

		res, err := controller.visualRepresentationInteractor.GetTestingResultByAccountID(ctx.Context(), id)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		if err := json.Unmarshal(res, &test); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		controller.coreInteractor.AddCognitiveSpecificationToUser(
			ctx.Context(),
			api_entity.User{
				Id: id,
			},
			api_entity.CognitiveSpecification{},
		)
		return ctx.Status(fiber.StatusAccepted).JSON(test)
	}
}

func (controller VisualRepresentationController) GetTestTemplateBySlug() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer controller.logger.CreateAPILog(ctx, time.Now())
		var (
			test entity.TestTemplate
		)
		slug := ctx.Params("slug")

		res, err := controller.visualRepresentationInteractor.GetTestTemplateBySlug(ctx.Context(), slug)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		if err := json.Unmarshal(res, &test); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusAccepted).JSON(test)
	}
}
