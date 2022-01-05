package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/habibiiberahim/go-backend/model"
)

func JWTProtected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("API_SECRET"),
		ErrorHandler: JwtErrorHandler,
	})
}

func JwtErrorHandler(c *fiber.Ctx, err error) error {

	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   nil,
		})

	} else {
		return c.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:   500,
			Status: "Internal Service Error",
			Data:   nil,
		})
	}

}
