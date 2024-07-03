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

func ListGifts(c *fiber.Ctx) error {
	req := new(Gift)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	gifts, err := ListAllGifts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve gifts",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"gifts": gifts,
	})
}

func AddGift(c *fiber.Ctx) error {
	req := new(Gift)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	err := AddGiftForUser(*req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to add gifts",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"GiftId": req.ID,
	})
}

func DeleteGift(c *fiber.Ctx) error {
	userId := c.Params("userId")
	giftId := c.Params("giftId")

	err := DeleteGiftForUser(userId, giftId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete gift",
		})
	}

	return c.SendString("Gift deleted successfully")
}
