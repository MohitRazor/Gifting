package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Use(logger.New()) // Logger middleware

	// Routes
	app.Get("/login", LoginUser)
	app.Get("/listGifts", ListGifts)
	app.Post("/addGift", AddGift)
	app.Delete("/deleteGift/:userId/:giftId", DeleteGift)
}
