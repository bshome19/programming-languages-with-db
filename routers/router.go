package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/bshome19/programming-languages-with-db/handlers"

)

type Router struct{}

var singleton *Router

func RouterInstance() *Router {
	if singleton == nil {
		singleton = &Router{}
	}
	return singleton
}

func (_ *Router) SetupRoutes(app *fiber.App) {

	app.Get("/languages", handlers.HandlerInstance().GetAllLanguages)
	app.Get("/languages/:id", handlers.HandlerInstance().GetLanguagesById)
	app.Post("/languages", handlers.HandlerInstance().CreateNewLanguage)
	app.Delete("/languages/:id", handlers.HandlerInstance().DeleteLanguageById)
	app.Put("/languages/:id", handlers.HandlerInstance().UpdateLanguageDataById)

}