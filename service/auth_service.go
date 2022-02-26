package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/go-backend/entity"
)

type AuthService interface {
	Redirect(c *fiber.Ctx) string
	Callback(c *fiber.Ctx) string
	CreateToken(user *entity.User) string
}
