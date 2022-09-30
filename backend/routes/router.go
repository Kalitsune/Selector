package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalitsune/selector/api"
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
	apiRoutes := app.Group("/api")
	apiRoutes.Use(handler.ApiMiddleware)

	apiRoutes.Get("/lists", handler.GetLists)
	apiRoutes.Post("/lists", handler.PostList)

	apiRoutes.Get("/list/:id", handler.GetList)
	apiRoutes.Patch("/list/:id", handler.PatchList)
	apiRoutes.Delete("/list/:id", handler.DeleteList)

	api.Logger.Info.Println("Routes initialized!")
}
