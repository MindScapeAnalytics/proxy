package visualrepresentation

import (
	"github.com/MindScapeAnalytics/proxy/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func VisualRepresentationGroup(mw middleware.MDWManager, accountRout fiber.Router, h VisualRepresentationController) {
	accountRout.Get("/assessment/account/:id", mw.NonAuthed(), h.GetTestingResultByAccountID())
	accountRout.Get("/questions/:slug", mw.NonAuthed(), h.GetTestTemplateBySlug())
}
