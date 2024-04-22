package middleware

import (
	"strings"

	api "github.com/MindScapeAnalytics/grpc-api/authentication/client"
	"github.com/MindScapeAnalytics/proxy/config"
	"github.com/MindScapeAnalytics/proxy/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type mdwManager struct {
	cfg                   *config.Config
	logger                logger.Logger
	authenticationService api.AuthenticationService
}

type MDWManager interface {
	APIMiddleware() fiber.Handler
	NonAuthed() fiber.Handler
}

func NewMDWManager(
	cfg *config.Config,
	logger logger.Logger,
	authenticationService api.AuthenticationService,
) MDWManager {
	return &mdwManager{
		cfg:                   cfg,
		logger:                logger,
		authenticationService: authenticationService,
	}
}

func (mw *mdwManager) APIMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		tokenSplit := strings.Split(tokenString, " ")
		if len(tokenSplit) == 2 {
			ok, err := mw.authenticationService.ValidateToken(c.Context(), []byte(tokenSplit[1]))
			if err != nil {
				return c.SendStatus(fiber.StatusUnauthorized)
			}
			if ok {
				return c.Next()
			}
		}
		return c.SendStatus(fiber.StatusUnauthorized)
	}
}

func (mw *mdwManager) NonAuthed() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
