package main

import (
	"log"

	"github.com/amirul-zafrin/snake-validator/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong!")
	})
	app.Get("/new", handlers.GenerateFruit)
	app.Post("/validate", handlers.Validate)
}
func main() {
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
