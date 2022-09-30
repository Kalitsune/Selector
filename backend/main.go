package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalitsune/selector/api"
	"github.com/kalitsune/selector/routes"
	"os"
)

func main() {
	// create a gofiber server
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	//define port
	port, portDefined := os.LookupEnv("PORT")
	if !portDefined {
		port = "8080"
	}

	//init api routes
	routes.Init(app)

	// start the server
	api.Logger.Info.Printf("Starting server on http://localhost:%s", port)
	api.Logger.Error.Fatalln(app.Listen(":" + port))
}
