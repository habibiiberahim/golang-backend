package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/go-backend/entity"
)

type AuthService interface {
	CheckPasswordHash(password, hash string)
	FindById(id string) (*entity.User, error)
	FindByName(name string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Login(c fiber.Ctx)
}
