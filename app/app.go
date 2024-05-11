package app

import (
	"context"

	"github.com/MindScapeAnalytics/proxy/config"
	accountIntr "github.com/MindScapeAnalytics/proxy/internal/controller/http/account"
	psychologyTestingIntr "github.com/MindScapeAnalytics/proxy/internal/controller/http/psychology_testing"
	visualRepresentationIntr "github.com/MindScapeAnalytics/proxy/internal/controller/http/visual_representation"
	accountRepo "github.com/MindScapeAnalytics/proxy/internal/interactor/account"
	psychologyTestingRepo "github.com/MindScapeAnalytics/proxy/internal/interactor/psychology_testing"
	visualRepresentationRepo "github.com/MindScapeAnalytics/proxy/internal/interactor/visual_representation"

	"github.com/MindScapeAnalytics/proxy/internal/middleware"
	"github.com/MindScapeAnalytics/proxy/pkg/transport/http"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	Fiber       *fiber.App
	Config      *config.Config
	Adapters    *Adapters
	Interactors *Interactors
	Controllers *Controllers
	Drivers     *Drivers
	Middleware  *Middleware
}

type Middleware struct {
	Middleware middleware.MDWManager
}

type Drivers struct {
	HTTPClient *http.Client
}

// Adapters ...
type Adapters struct {
	AccountRepository        accountRepo.AccountRepository
	VisualRepresentationRepo visualRepresentationRepo.VisualRepresentationRepository
	PsychologyTestingRepo    psychologyTestingRepo.PsychologyTestingRepository
}

// Interactors ...
type Interactors struct {
	AccountInteractor              accountIntr.AccountInteractor
	PsychologyTestingInteractor    psychologyTestingIntr.PsychologyTestingInteractor
	VisualRepresentationInteractor visualRepresentationIntr.VisualRepresentationInteractor
}

// Controllers ...
type Controllers struct {
	HTTP struct {
		AccountController              accountIntr.AccountController
		PsychologyTestingController    psychologyTestingIntr.PsychologyTestingController
		VisualRepresentationController visualRepresentationIntr.VisualRepresentationController
	}
}

func Run(cfg *config.Config, ctx context.Context) error {
	app, err := newApp(ctx, cfg)
	if err != nil {
		return err
	}

	err = app.Drivers.HTTPClient.Launch()
	if err != nil {
		return err
	}

	return nil
}
