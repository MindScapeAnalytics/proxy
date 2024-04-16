package logger

import (
	"fmt"
	"time"

	"github.com/MindScapeAnalytics/proxy/config"
	"github.com/MindScapeAnalytics/proxy/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type loggerUC struct {
	cfg    *config.Config
	logger Logger
}

type LoggerUC interface {
	CreateAPILog(c *fiber.Ctx, startAt time.Time)
}

func NewLoggerUC(
	cfg *config.Config,
	logger Logger,
) LoggerUC {
	return &loggerUC{
		logger: logger,
		cfg:    cfg,
	}
}

func (l *loggerUC) CreateAPILog(c *fiber.Ctx, startAt time.Time) {
	apiLogData := APILogData{
		IP:       c.IP(),
		Endpoint: fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()),
		//RequestBody:  string(c.Body()),
		//ResponseBody: string(c.Response().Body()),
		StatusCode: c.Response().StatusCode(),
		Took:       time.Since(startAt).Milliseconds(),
		StartAt:    startAt,
		CreatedAt:  time.Now(),
	}
	if errorMessage, ok := c.Locals("error").(string); ok {
		apiLogData.ErrMessage = &errorMessage
		l.logger.Warnf("%s", *apiLogData.ErrMessage)
	}
	l.logger.Debugf("%s", utils.GetStructJSON(apiLogData))
}
