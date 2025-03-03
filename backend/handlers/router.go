package handlers

import (
	"backend/middlewares"
	authenticaion "backend/service/auth"
	"backend/service/testservice"
	"backend/service/users"

	// "backend/service/users"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	authEndpoint := authenticaion.InitAuthenEndpoint()
	testserviceEndpoint := testservice.InitTestServiceEndpoint()
	userserviceEndpoint := users.InitUserEndpoint()

	// Public Routes
	app.Post("/login", authEndpoint.Login)
	app.Post("/register", authEndpoint.Register)
	app.Get("/logout", authEndpoint.Logout)

	// use Google middleware only login with google
	google := app.Group("/google", middlewares.SetverifyGoogleTokenMiddleware)
	google.Post("/login", authEndpoint.LoginWithGoogle)
	google.Post("/register", authEndpoint.RegisterWithGoogle)

	// Private Routes (Protected by JWT)
	private := app.Group("/api", middlewares.SetJWTHandler())
	private.Get("/test", testserviceEndpoint.TestService)

	// User
	users := app.Group("/user", middlewares.SetJWTHandler())
	// users.Post("/create", userserviceEndpoint.CreateUser)
	users.Get("/", userserviceEndpoint.GetUser)
}
