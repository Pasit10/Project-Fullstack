package users

import (
	templateError "backend/error"

	"github.com/gofiber/fiber/v2"
)

type userEndpoint struct {
	userLogic *userLogic
}

func InitUserEndpoint() *userEndpoint {
	logic := &userLogic{}
	return &userEndpoint{
		userLogic: logic.InitUserLogic(),
	}
}

func (ue *userEndpoint) CreateUser(c *fiber.Ctx) error {
	var req User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := ue.userLogic.CreateUserLogic(req); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.BadrequestError)
		return c.Status(httpStatusCode).JSON(errorResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func (ue *userEndpoint) GetUser(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)

	user_res, err := ue.userLogic.GetuserData(uid)
	if err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		return c.Status(httpStatusCode).JSON(errorResponse)
	}
	return c.Status(fiber.StatusOK).JSON(user_res)
}
