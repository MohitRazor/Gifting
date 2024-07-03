package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Use(logger.New()) // Logger middleware

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/login", LoginUser)
	app.Post("/register", RegisterUser)
	app.Get("/gifts", GetAllGifts)
	app.Get("/gifts/:email", GetGiftsForUser)
	app.Post("/gifts/:email", AddGiftForUser)
	app.Post("/gifts", AddGift)
	app.Delete("/gift/:email", DeleteGift)
}
