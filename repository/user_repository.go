package repository

import (
	"github.com/habibiiberahim/go-backend/entity"
)

type UserRepository interface {
	FindAll() []entity.User
}
