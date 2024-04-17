package http

import (
	"errors"
	"fmt"
	"log"

	"github.com/MindScapeAnalytics/proxy/config"
	"github.com/MindScapeAnalytics/proxy/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
)

type Client struct {
	app *fiber.App
	cfg *config.Config
}

func NewClientHTTP(cfg *config.Config, logger logger.Logger) (Client, error) {
	app := fiber.New()
	return Client{
		app: app,
		cfg: cfg,
	}, nil
}

func (c *Client) Launch() error {
	if c.app == nil {
		return errors.New("Серевер не проинициализирован ")
	}

	c.app.Use(fiberLogger.New())
	c.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	log.Default().Printf("Сервер запущен на: %s", c.cfg.Server.Port)
	if err := c.app.Listen(fmt.Sprintf(":%s", c.cfg.Server.Port)); err != nil {
		log.Default().Printf("Ошибка запуска сервера: ", err)
		return err
	}
	return nil
}

func (c *Client) Shutdown() error {
	err := c.app.Shutdown()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetApp() (*fiber.App, error) {
	if c.app != nil {
		return c.app, nil
	} else {
		return nil, errors.New("empty fiber app")
	}
}
