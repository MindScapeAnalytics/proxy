package app

import (
	"context"
	"log"

	authenticationService "github.com/MindScapeAnalytics/grpc-api/authentication/client"
	coreService "github.com/MindScapeAnalytics/grpc-api/core/client"
	psychologyTestingService "github.com/MindScapeAnalytics/grpc-api/psychology_testing/client"
	visualRepresentationService "github.com/MindScapeAnalytics/grpc-api/visual_representation/client"
	"github.com/MindScapeAnalytics/proxy/config"
	accountRepo "github.com/MindScapeAnalytics/proxy/internal/adapters/account"
	core "github.com/MindScapeAnalytics/proxy/internal/adapters/core"
	psychologytesting "github.com/MindScapeAnalytics/proxy/internal/adapters/psychology_testing"
	visualrepresentation "github.com/MindScapeAnalytics/proxy/internal/adapters/visual_representation"
	accountCtrl "github.com/MindScapeAnalytics/proxy/internal/controller/http/account"
	coreIntr "github.com/MindScapeAnalytics/proxy/internal/controller/http/core"
	psychologyTestingIntr "github.com/MindScapeAnalytics/proxy/internal/controller/http/psychology_testing"
	visualRepresentationIntr "github.com/MindScapeAnalytics/proxy/internal/controller/http/visual_representation"
	accountIntr "github.com/MindScapeAnalytics/proxy/internal/interactor/account"
	coreRepo "github.com/MindScapeAnalytics/proxy/internal/interactor/core"
	psychologytestingRepo "github.com/MindScapeAnalytics/proxy/internal/interactor/psychology_testing"
	visualrepresentationRepo "github.com/MindScapeAnalytics/proxy/internal/interactor/visual_representation"
	"github.com/MindScapeAnalytics/proxy/internal/middleware"
	"github.com/MindScapeAnalytics/proxy/pkg/logger"
	"github.com/MindScapeAnalytics/proxy/pkg/transport/http"
)

func newApp(ctx context.Context, cfg *config.Config) (*App, error) {

	app := &App{}

	l := logger.NewLogger(cfg)
	err := l.InitLogger()
	if err != nil {
		return nil, err
	}
	logger := logger.NewLoggerUC(cfg, l)

	httpClient, err := http.NewClientHTTP(cfg, l)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	authenticationService, err := authenticationService.NewAuthenticationService(ctx, cfg.AuthenticationService.IP, cfg.AuthenticationService.Port)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	visualRepresentationService, err := visualRepresentationService.NewVisualRepresentationService(ctx, cfg.VisualRepresentationService.IP, cfg.VisualRepresentationService.Port)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	psychologyTestingService, err := psychologyTestingService.NewCoreService(ctx, cfg.PsychologyTestingService.IP, cfg.PsychologyTestingService.Port)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	coreService, err := coreService.NewCoreService(ctx, cfg.CoreService.IP, cfg.CoreService.Port)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	if app.Fiber, err = httpClient.GetApp(); err != nil {
		return nil, err
	}

	app.Drivers = &Drivers{
		HTTPClient: &httpClient,
	}

	err = app.initAdapters(
		ctx,
		cfg,
		authenticationService,
		visualRepresentationService,
		psychologyTestingService,
		coreService,
	)
	if err != nil {
		return nil, err
	}
	err = app.initInteractors(ctx)
	if err != nil {
		return nil, err
	}

	if err = app.initMiddleware(ctx, cfg, l, authenticationService); err != nil {
		return nil, err
	}

	err = app.initControllers(ctx, logger)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) initAdapters(
	ctx context.Context,
	cfg *config.Config,
	authenticationService authenticationService.AuthenticationService,
	visualRepresentationService visualRepresentationService.VisualRepresentationService,
	psychologyTestingService psychologyTestingService.PsychologyTestingService,
	coreService coreService.CoreService,
) (err error) {

	adapters := &Adapters{}

	// инициализируем user repository
	if adapters.AccountRepository, err = accountRepo.NewAccountRepository(ctx, accountRepo.AccountRepOpts{
		AuthenticationService: authenticationService,
		Type:                  "authenticationService",
	}); err != nil {
		return err
	}

	if adapters.VisualRepresentationRepo, err = visualrepresentation.NewVisualRepresentation(ctx, visualrepresentation.VisualRepresentationRepOpts{
		VisualRepresentationService: visualRepresentationService,
		Type:                        "visualRepresentationService",
	}); err != nil {
		return err
	}

	if adapters.PsychologyTestingRepo, err = psychologytesting.NewPsychologyTestingRepository(ctx, psychologytesting.PsychologyTestingRepositoryOpts{
		PsychologyTestingService: psychologyTestingService,
		Type:                     "psychologyTestingService",
	}); err != nil {
		return err
	}

	if adapters.CoreRepository, err = core.NewCoreRepository(ctx, core.CoreRepositoryOpts{
		CoreService: coreService,
		Type:        "coreService",
	}); err != nil {
		return err
	}

	app.Adapters = adapters

	return nil
}

func (app *App) initInteractors(ctx context.Context) (err error) {

	interactors := &Interactors{}

	if interactors.AccountInteractor, err = accountIntr.NewAccountInteractor(ctx, accountIntr.AccountIntrOpts{
		CoreRepository:    app.Adapters.CoreRepository,
		AccountRepository: app.Adapters.AccountRepository,
	}); err != nil {
		return err
	}
	if interactors.PsychologyTestingInteractor, err = psychologytestingRepo.NewPsychologyTestingInteractor(ctx, psychologytestingRepo.PsychologyTestingInteractorOpts{
		PsychologyTestingRepository: app.Adapters.PsychologyTestingRepo,
	}); err != nil {
		return err
	}
	if interactors.VisualRepresentationInteractor, err = visualrepresentationRepo.NewVisualRepresentationInteractor(ctx, visualrepresentationRepo.VisualRepresentationInteractorOpts{
		VisualRepresentationRepository: app.Adapters.VisualRepresentationRepo,
	}); err != nil {
		return err
	}
	if interactors.CoreInteractor, err = coreRepo.NewCoreInteractor(ctx, coreRepo.CoreInteractorOpts{
		VisualRepresentation: app.Adapters.VisualRepresentationRepo,
		CoreRepository:       app.Adapters.CoreRepository,
	}); err != nil {
		return err
	}

	app.Interactors = interactors

	return nil
}

func (app *App) initMiddleware(ctx context.Context, cfg *config.Config, logger logger.Logger, authenticationService authenticationService.AuthenticationService) (err error) {
	app.Middleware = &Middleware{}
	app.Middleware.Middleware = middleware.NewMDWManager(cfg, logger, authenticationService)
	return nil
}

func (app *App) initControllers(ctx context.Context, logger logger.LoggerUC) (err error) {

	controllers := &Controllers{}

	// инициализируем user http controller
	if controllers.HTTP.AccountController, err = accountCtrl.NewAccountController(ctx, accountCtrl.AccountCtrlOpts{
		AccountInteractor: app.Interactors.AccountInteractor,
		Logger:            logger,
	}); err != nil {
		return err
	}
	if controllers.HTTP.VisualRepresentationController, err = visualRepresentationIntr.NewVisualRepresentationController(ctx, visualRepresentationIntr.VisualRepresentationControllerOpts{
		CoreInteractor:                 app.Interactors.CoreInteractor,
		VisualRepresentationInteractor: app.Interactors.VisualRepresentationInteractor,
		Logger:                         logger,
	}); err != nil {
		return err
	}

	if controllers.HTTP.PsychologyTestingController, err = psychologyTestingIntr.NewPsychologyTestingController(ctx, psychologyTestingIntr.PsychologyTestingCtrlOpts{
		PsychologyTestingInteractor: app.Interactors.PsychologyTestingInteractor,
		Logger:                      logger,
	}); err != nil {
		return err
	}

	if controllers.HTTP.CoreController, err = coreIntr.NewCoreController(ctx, coreIntr.CoreControllerOpts{
		CoreInteractor: app.Interactors.CoreInteractor,
		Logger:         logger,
	}); err != nil {
		return err
	}

	app.Controllers = controllers

	api := app.Fiber.Group("/api/v1")
	accountRouter := api.Group("/account")
	testing := api.Group("/testing")
	core := api.Group("/core")

	accountCtrl.AccountRoutesGroup(app.Middleware.Middleware, accountRouter, controllers.HTTP.AccountController)
	visualRepresentationIntr.VisualRepresentationGroup(app.Middleware.Middleware, testing, controllers.HTTP.VisualRepresentationController)
	psychologyTestingIntr.PsychologyTestingGroup(app.Middleware.Middleware, testing, controllers.HTTP.PsychologyTestingController)
	coreIntr.CoreRoutesGroup(app.Middleware.Middleware, core, controllers.HTTP.CoreController)
	return nil
}
