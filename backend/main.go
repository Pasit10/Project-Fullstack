package main

import (
	"backend/internal/gateways"
	repo "backend/internal/repositories"
	ser "backend/internal/services"
	"backend/middlewares"
	"backend/pkg/database"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("err load ENV")
		os.Exit(1)
	}

	app := fiber.New()
	app.Use(middlewares.NewLogger())
	// app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://localhost:5173",
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
	}))

	// SetUp Database
	mysql := database.InitDatabaseConnection()
	// config.InitFirebase()

	authRepo := repo.InitAuthRepository(mysql)

	authService := ser.InitAuthenService(authRepo)

	gateways.InitHTTPGateway(app, authService)

	PORT := os.Getenv("SERVER_PORT")
	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
