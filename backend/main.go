package main

import (
	"backend/config"
	"backend/handlers"
	"backend/middlewares"
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

	// SetUp Database
	config.InitDatabaseConnection()
	config.InitFirebase()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://localhost:5173",
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
	}))

	app.Use(middlewares.NewLogger())

	handlers.InitRoutes(app)

	PORT := os.Getenv("SERVER_PORT")
	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
