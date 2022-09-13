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

	//API routes
	api := app.Group("/api")

	//auth routes
	auth := api.Group("/auth")
	auth.Get("/", handler.GoogleAuth)
	auth.Get("/callback", handler.GoogleCallback)
}
