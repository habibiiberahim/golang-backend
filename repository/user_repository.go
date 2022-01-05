package repository

import (
	"github.com/habibiiberahim/go-backend/entity"
)

type UserRepository interface {
	GetAll() []entity.User
}
