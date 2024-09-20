package routes

import (
	"user-api/controllers"
	"user-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)

	app.Post("/api/refresh-token", controllers.RefreshToken)

	app.Get("/api/user", middlewares.AuthMiddleware(), controllers.GetUserDetail)
}
