package gateways

import (
	"backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

type HTTPGateway struct {
	AuthService services.IAuthService
}

func InitHTTPGateway(app *fiber.App, auth services.IAuthService) {
	gateway := &HTTPGateway{
		AuthService: auth,
	}

	InitRoutes(app, *gateway)
}
