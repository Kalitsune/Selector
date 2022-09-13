package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalitsune/selector/routes"
	"log"
)

func main() {
	// create a gofiber server
	app := fiber.New()

	routes.Init(app)

	log.Fatalln(app.Listen(":8080"))
}
