package middleware

import (
	"github.com/Simonz2/Task_app/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {

	// Check for the presence of a token in the Authorization header
	tokenString := c.Get("Authorization")
	if tokenString == "" || len(tokenString) < 7 || tokenString[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing or invalid token"})
	}
	tokenString = tokenString[7:] // Remove "Bearer " prefix
	// Validate the token
	if err := utils.VerifyToken(tokenString); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
	}
	// If the token is valid, proceed to the next handler
	return c.Next()

}
