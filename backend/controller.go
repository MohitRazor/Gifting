package main

import "github.com/gofiber/fiber/v2"

func LoginUser(c *fiber.Ctx) error {
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	user, err := FindUser(req.Username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})

	}
	if user.Password != req.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged in",
	})
}
