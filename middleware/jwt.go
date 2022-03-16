package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/habibiiberahim/go-backend/config"
	"github.com/habibiiberahim/go-backend/model"
)

func Protected() fiber.Handler {
	configuration := config.NewConfiguration()
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(configuration.Get("APP_SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(
				model.WebResponse{
					Code:   400,
					Status: "Error",
					Data:   "Missing or malformed JWT",
				})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(
			model.WebResponse{
				Code:   401,
				Status: "Error",
				Data:   "Invalid or expired JWT",
			})
}
