package middlewares

import (
	"context"
	"strings"

	"backend/config" // Update this with your actual project path

	"github.com/gofiber/fiber/v2"
)

// FirebaseAuthMiddleware validates Firebase JWT token
func FirebaseAuthMiddleware(c *fiber.Ctx) error {
	firebaseAuth := config.FB.AuthDatabase
	if firebaseAuth == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Firebase not initialized"})
	}

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := firebaseAuth.VerifyIDToken(context.Background(), tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}
	c.Locals("user", token)

	return c.Next()
}
