package controller

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	Callback(c *fiber.Ctx) error
}
