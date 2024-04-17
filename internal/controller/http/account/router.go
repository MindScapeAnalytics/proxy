package user

import (
	"github.com/MindScapeAnalytics/proxy/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func AccountRoutesGroup(mw middleware.MDWManager, accountRout fiber.Router, h AccountController) {
	accountRout.Post("/", mw.NonAuthed(), h.Login())
	accountRout.Post("/token", mw.NonAuthed(), h.Registry())
}
