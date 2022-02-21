package repository

import (
	"github.com/habibiiberahim/go-backend/entity"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepositoryImpl{
		DB: db,
	}
}

type userRepositoryImpl struct {
	DB *gorm.DB
}

func (repository userRepositoryImpl) FindAll() (users []entity.User) {
	return users
}
