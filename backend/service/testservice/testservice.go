package testservice

import "github.com/gofiber/fiber/v2"

type TestserviceEndpoint struct{}

func InitTestServiceEndpoint() *TestserviceEndpoint {
	return &TestserviceEndpoint{}
}

func (te *TestserviceEndpoint) TestService(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Test service successful"})
}
