package middlewares

import (
	"user-api/auth"
	"user-api/utils"

	"github.com/gofiber/fiber/v2"
)

// Middleware to check if the token is blacklisted
func TokenBlacklistMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(utils.ErrTokenNotProvided.Code).JSON(utils.ErrTokenNotProvided)
		}

		if auth.IsTokenBlacklisted(token) {
			return c.Status(utils.ErrTokenBlacklisted.Code).JSON(utils.ErrTokenBlacklisted)
		}

		return c.Next()
	}
}
