package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalitsune/selector/handler"
)

func Init(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		//TODO: load the index.html file
		return c.SendString("Hello, World!")
	})

	//auth routes
	auth := app.Group("/auth")
	auth.Get("/", handler.GoogleAuth)
	auth.Get("/callback", handler.GoogleCallback)

	//API routes
	api := app.Group("/api")
	api.Get("/lists", handler.GetLists)
}
