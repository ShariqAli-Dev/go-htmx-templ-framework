package api

import (
	"github.com/gofiber/fiber/v2"
)

func ValidateUser(c *fiber.Ctx) error {
	_, err := GetUserFromSession(c)
	if err != nil {
		return c.Redirect("/")
	}
	return c.Next()
}
