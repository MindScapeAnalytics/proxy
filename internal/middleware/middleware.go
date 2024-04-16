package middleware

import (
	"github.com/MindScapeAnalytics/proxy/config"
	"github.com/MindScapeAnalytics/proxy/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type mdwManager struct {
	cfg    *config.Config
	logger logger.Logger
}

type MDWManager interface {
	APIMiddleware() fiber.Handler
	NonAuthed() fiber.Handler
}

func NewMDWManager(
	cfg *config.Config,
	logger logger.Logger,
) MDWManager {
	return &mdwManager{
		cfg:    cfg,
		logger: logger,
	}
}

func (mw *mdwManager) APIMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// if c.Get("API-Key") != mw.cfg.Server.APIKey {
		// 	mw.logger.Warn(" APIMiddleware wrong API Key")
		// 	return c.SendStatus(fiber.StatusUnauthorized)
		// }
		// sessionKey := c.Get("Authorization")
		// authed, err := utils.ValidateSession(mw.db, sessionKey)
		// if err != nil {
		// 	return c.SendStatus(fiber.StatusUnauthorized)
		// }
		// if authed {
		// 	return c.Next()
		// } else {
		// 	return c.SendStatus(fiber.StatusUnauthorized)
		// }
		return c.Next()
	}
}

func (mw *mdwManager) NonAuthed() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//if c.Get("API-Key") != mw.cfg.Server.APIKey {
		//	mw.logger.Warn(" APIMiddleware wrong API Key")
		//	return c.SendStatus(fiber.StatusUnauthorized)
		//}
		return c.Next()
	}
}
