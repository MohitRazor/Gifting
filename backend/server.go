package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {

	// Connect to the database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")

	// Initialize a new Fiber app
	app := fiber.New()
	app.Use(cors.New())
	SetupRoutes(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3001"))
}
