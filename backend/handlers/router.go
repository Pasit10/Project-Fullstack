package handlers

import (
	"backend/middlewares"
	"backend/service/testservice"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	// authEndpoint := authenticaion.InitAuthenEndpoint()
	testserviceEndpoint := testservice.InitTestServiceEndpoint()

	// Public Routes
	// app.Post("/login", authEndpoint.Login)
	// app.Post("/register", authEndpoint.Register)
	// app.Get("/logout", authEndpoint.Logout)

	// Private Routes (Protected by JWT)
	private := app.Group("/api", middlewares.FirebaseAuthMiddleware)
	private.Get("/test", testserviceEndpoint.TestService)
}
