package main

import (
	"github.com/gofiber/fiber/v2"
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

	SetupRoutes(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
