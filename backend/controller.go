package main

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parser JSON",
		})
	}
	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}
	existingUser, err := FindUser(user.Email)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Cannot validate Email: " + err.Error(),
			})
		}
	}
	if existingUser.Email != "" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email already exists",
		})
	}
	user.Gifts = []string{}
	if err := InsertUser(*user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot insert user: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created",
	})
}

func LoginUser(c *fiber.Ctx) error {
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	user, err := FindUser(req.Email)
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

func GetGiftsForUser(c *fiber.Ctx) error {
	email := c.Params("email")
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email is required",
		})
	}
	giftIds, err := GetGiftIdsForUser(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	gifts, err := GetGiftByIds(giftIds)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get gifts",
		})
	}
	return c.Status(fiber.StatusOK).JSON(gifts)
}

func AddGift(c *fiber.Ctx) error {
	req := new(Gift)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if req.Name == "" || req.Price == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name and price are required",
		})
	}
	if err := InsertGift(*req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot insert gift",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Gift created",
	})
}

func AddGiftForUser(c *fiber.Ctx) error {
	email := c.Params("email")
	giftId := c.Query("giftId")
	if email == "" || giftId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username and gift ID are required",
		})
	}
	gift, err := GetGiftById(giftId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Gift not found",
		})
	}
	// Check if Gift is already added
	giftIds, err := GetGiftIdsForUser(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Error getting gifts for user",
		})
	}
	for _, id := range giftIds {
		if id == giftId {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Gift already added",
			})
		}
	}

	if err := InsertGiftForUser(email, gift.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot add gift to user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Gift added to user",
	})
}

func GetAllGifts(c *fiber.Ctx) error {
	gifts, err := ListAllGifts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get gifts",
		})
	}
	return c.Status(fiber.StatusOK).JSON(gifts)
}

func DeleteGift(c *fiber.Ctx) error {
	email := c.Params("email")
	giftId := c.Query("giftId")
	if email == "" || giftId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username and gift ID are required",
		})
	}
	giftIds, err := GetGiftIdsForUser(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	found := false
	for _, id := range giftIds {
		if id == giftId {
			found = true
			break
		}
	}
	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Gift not found",
		})
	}
	if err := DeleteGiftForUser(email, giftId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete gift",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Gift deleted",
	})
}
