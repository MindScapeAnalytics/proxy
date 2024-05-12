package user

import (
	"github.com/MindScapeAnalytics/proxy/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func CoreRoutesGroup(mw middleware.MDWManager, accountRout fiber.Router, h CoreController) {
	accountRout.Post("/token", mw.APIMiddleware(), h.AddUser())
	accountRout.Post("/", mw.APIMiddleware(), h.AddEvent())
	accountRout.Post("/token", mw.APIMiddleware(), h.AddFriend())
	accountRout.Post("/", mw.APIMiddleware(), h.RemoveFriend())
	accountRout.Post("/token", mw.APIMiddleware(), h.AddEventToUser())
	accountRout.Post("/", mw.APIMiddleware(), h.AddCognitiveSpecificationToUser())
	accountRout.Post("/token", mw.APIMiddleware(), h.UpdateCognitiveSpecificationUser())
	accountRout.Post("/", mw.APIMiddleware(), h.UpdateAdditionalUserEventInfo())
	accountRout.Post("/token", mw.APIMiddleware(), h.GetUserEventList())
}
