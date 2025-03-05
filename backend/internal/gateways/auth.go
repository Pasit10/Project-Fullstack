package gateways

import (
	templateError "backend/error"
	"backend/internal/entities"
	"backend/middlewares"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h HTTPGateway) Login(c *fiber.Ctx) error {
	var req entities.UserLogin
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	isValid, user, err := h.AuthService.Login(req)
	if err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		return c.Status(httpStatusCode).JSON(errorResponse)
	}

	if !isValid {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.UnauthorizedError)
		return c.Status(httpStatusCode).JSON(errorResponse)
	}

	token, err := middlewares.GenerateJWT(user.UID, user.Name, user.Role) ///TODO: change role
	if err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		return c.Status(httpStatusCode).JSON(errorResponse)
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		SameSite: "None",
		Secure:   false, // Use true if on HTTPS
	})
	return c.JSON(fiber.Map{"message": "Login successful"})
}

func (h HTTPGateway) Register(c *fiber.Ctx) error {
	var req entities.User
	req.UID = uuid.New().String()
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	req.Role = "user"
	err := h.AuthService.Register(req)
	if err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		return c.Status(httpStatusCode).JSON(errorResponse)
	}

	token, err := middlewares.GenerateJWT(req.UID, req.Name, req.Role) ///TODO: change role
	if err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		return c.Status(httpStatusCode).JSON(errorResponse)
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		SameSite: "None",
		Secure:   false,
	})
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created"})
}

func (h HTTPGateway) Logout(c *fiber.Ctx) error {
	// Clear the JWT cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		HTTPOnly: true,
		Expires:  time.Now().Add(-1 * time.Hour), // Expire immediately
	})
	return c.JSON(fiber.Map{"message": "Logout successful"})
}

func (h HTTPGateway) RegisterWithGoogle(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)
	email := c.Locals("email").(string)
	name := c.Locals("name").(string)
	picture := c.Locals("picture").(string)

	var req = entities.User{
		UID:     uid,
		Email:   email,
		Name:    name,
		Picture: picture,
		Role:    "user",
	}
	req.Role = "user"
	err := h.AuthService.RegisterGoogle(req)
	if err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		return c.Status(httpStatusCode).JSON(errorResponse)
	}

	token, err := middlewares.GenerateJWT(req.UID, req.Name, req.Role) ///TODO: change role
	if err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		return c.Status(httpStatusCode).JSON(errorResponse)
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		SameSite: "None",
		Secure:   false,
	})
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created"})
}

func (h HTTPGateway) LoginWithGoogle(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)
	user, err := h.AuthService.GetUserByUID(uid)
	if err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		return c.Status(httpStatusCode).JSON(errorResponse)
	}

	token, err := middlewares.GenerateJWT(user.UID, user.Name, user.Role) ///TODO: change role
	if err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		return c.Status(httpStatusCode).JSON(errorResponse)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		SameSite: "None",
		Secure:   false, // Use true if on HTTPS
	})
	return c.JSON(fiber.Map{"message": "Login successful"})
}
