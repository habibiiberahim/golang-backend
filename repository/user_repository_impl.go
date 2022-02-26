package repository

import (
	"github.com/danilopolani/gocialite/structs"
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
func (repository userRepositoryImpl) Store() {

}

func (repository userRepositoryImpl) GetOrRegister(provider string, user *structs.User) entity.User {
	userData := entity.User{
		FullName: user.FullName,
		Email:    user.Email,
		SocialId: user.ID,
		Provider: provider,
	}

	repository.DB.Where("provider = ? AND social_id = ? ", provider, user.ID).FirstOrCreate(&userData)

	return userData

}
