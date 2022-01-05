package service

import (
	"github.com/habibiiberahim/go-backend/model"
	"github.com/habibiiberahim/go-backend/repository"
)

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
	}
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func (service *userServiceImpl) List() (responses []model.GetProductResponse) {
	users := service.UserRepository.GetAll()
	for _, user := range users {
		responses = append(responses, model.GetProductResponse{
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return responses
}
