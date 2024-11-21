package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tunde262/trivia_rest_api_docker_postgres_go/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}