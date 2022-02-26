package repository

import (
	"github.com/danilopolani/gocialite/structs"
	"github.com/habibiiberahim/go-backend/entity"
)

type UserRepository interface {
	GetOrRegister(provider string, user *structs.User) entity.User
}
