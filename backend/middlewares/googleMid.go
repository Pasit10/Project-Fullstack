package middlewares

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/idtoken"
)

var req struct {
	Token string `json:"token"`
}

func SetverifyGoogleTokenMiddleware(c *fiber.Ctx) error {
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Token is required"})
	}

	payload, err := idtoken.Validate(context.Background(), req.Token, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		log.Println("Token verification failed:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	uid := payload.Claims["sub"].(string)         // User ID
	email := payload.Claims["email"].(string)     // Email
	name := payload.Claims["name"].(string)       // Full Name
	picture := payload.Claims["picture"].(string) // Profile Picture

	c.Locals("uid", uid)
	c.Locals("email", email)
	c.Locals("name", name)
	c.Locals("picture", picture)

	return c.Next()
}
