package psychologytesting

import (
	"github.com/MindScapeAnalytics/proxy/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func PsychologyTestingGroup(mw middleware.MDWManager, accountRout fiber.Router, h PsychologyTestingController) {
	accountRout.Post("/answers", mw.APIMiddleware(), h.SendTestingData())
}
