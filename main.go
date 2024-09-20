package main

import (
	"log"
	"user-api/database"
	"user-api/middlewares"
	"user-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	database.Migrate()

	app := fiber.New()

	middlewares.SetupCORS(app)

	routes.SetUpRoutes(app)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
