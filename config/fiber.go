package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/go-backend/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
