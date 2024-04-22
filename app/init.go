package app

import (
	"context"
	"log"

	api "github.com/MindScapeAnalytics/grpc-api/authentication/client"
	"github.com/MindScapeAnalytics/proxy/config"
	accountRepo "github.com/MindScapeAnalytics/proxy/internal/adapters/account"
	accountCtrl "github.com/MindScapeAnalytics/proxy/internal/controller/http/account"
	accountIntr "github.com/MindScapeAnalytics/proxy/internal/interactor/account"
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

	authenticationService, err := api.NewAuthenticationService(ctx, cfg.AuthenticationService.IP, cfg.AuthenticationService.Port)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	if app.Fiber, err = httpClient.GetApp(); err != nil {
		return nil, err
	}

	app.Drivers = &Drivers{
		HTTPClient: &httpClient,
	}

	err = app.initAdapters(ctx, cfg, authenticationService)
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

func (app *App) initAdapters(ctx context.Context, cfg *config.Config, authenticationService api.AuthenticationService) (err error) {

	adapters := &Adapters{}

	// инициализируем user repository
	if adapters.AccountRepository, err = accountRepo.NewAccountRepository(ctx, accountRepo.AccountRepOpts{
		AuthenticationService: authenticationService,
		Type:                  "authenticationService",
	}); err != nil {
		return err
	}

	app.Adapters = adapters

	return nil
}

func (app *App) initInteractors(ctx context.Context) (err error) {

	interactors := &Interactors{}

	// инициализируем user interactor
	if interactors.AccountInteractor, err = accountIntr.NewAccountInteractor(ctx, accountIntr.AccountIntrOpts{
		// прокидываем настройки
		AccountRepository: app.Adapters.AccountRepository,
	}); err != nil {
		return err
	}

	app.Interactors = interactors

	return nil
}

func (app *App) initMiddleware(ctx context.Context, cfg *config.Config, logger logger.Logger, authenticationService api.AuthenticationService) (err error) {
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

	app.Controllers = controllers

	api := app.Fiber.Group("/api/v1")
	accountRouter := api.Group("/account")

	accountCtrl.AccountRoutesGroup(app.Middleware.Middleware, accountRouter, controllers.HTTP.AccountController)

	return nil
}
