package user

import (
	"github.com/MindScapeAnalytics/proxy/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func CoreRoutesGroup(mw middleware.MDWManager, accountRout fiber.Router, h CoreController) {
	// accountRout.Post("/token", mw.APIMiddleware(), h.AddUser())
	accountRout.Post("/events", mw.APIMiddleware(), h.AddEvent())
	// accountRout.Post("/token", mw.APIMiddleware(), h.AddFriend())
	// accountRout.Post("/", mw.APIMiddleware(), h.RemoveFriend())
	// accountRout.Post("/token", mw.APIMiddleware(), h.AddEventToUser())
	// accountRout.Post("/add_cognitive", mw.APIMiddleware(), h.AddCognitiveSpecificationToUser())
	// accountRout.Post("/token", mw.APIMiddleware(), h.UpdateCognitiveSpecificationUser())
	// accountRout.Post("/", mw.APIMiddleware(), h.UpdateAdditionalUserEventInfo())
	accountRout.Get("/events", mw.APIMiddleware(), h.GetUserEventList()) //add limit
	accountRout.Post("/events", mw.APIMiddleware(), h.AddEventInfo())
	accountRout.Get("/events/:id", mw.APIMiddleware(), h.GetEventInfo())
}
