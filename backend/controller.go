package main

import "github.com/gofiber/fiber/v2"

func LoginUser(c *fiber.Ctx) error {
	return c.SendString("Login User")
}