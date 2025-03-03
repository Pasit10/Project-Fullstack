package authenticaion

import (
	templateError "backend/error"
	"backend/middlewares"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthenEndpoint struct {
	AuthenLogic *AuthenLogic
}

func InitAuthenEndpoint() *AuthenEndpoint {
	logic := &AuthenLogic{}
	return &AuthenEndpoint{
		AuthenLogic: logic.InitAuthenLogic(),
	}
}

func (ae *AuthenEndpoint) Login(c *fiber.Ctx) error {
	var req UserLogin
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	isValid, user, err := ae.AuthenLogic.LoginLogic(req)
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

func (ae *AuthenEndpoint) Register(c *fiber.Ctx) error {
	var req User
	req.UID = uuid.New().String()
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	req.Role = "user"
	err := ae.AuthenLogic.RegisterLogic(req)
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

func (ae *AuthenEndpoint) Logout(c *fiber.Ctx) error {
	// Clear the JWT cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		HTTPOnly: true,
		Expires:  time.Now().Add(-1 * time.Hour), // Expire immediately
	})
	return c.JSON(fiber.Map{"message": "Logout successful"})
}

func (ae *AuthenEndpoint) RegisterWithGoogle(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)
	email := c.Locals("email").(string)
	name := c.Locals("name").(string)
	picture := c.Locals("picture").(string)

	var req = User{
		UID:     uid,
		Email:   email,
		Name:    name,
		Picture: picture,
		Role:    "user",
	}
	req.Role = "user"
	err := ae.AuthenLogic.RegisterGoogleLogic(req)
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

func (ae *AuthenEndpoint) LoginWithGoogle(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)
	user, err := ae.AuthenLogic.GetUserByUID(uid)
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
