package service

import (
	"github.com/habibiiberahim/go-backend/model"
	"github.com/habibiiberahim/go-backend/repository"
)

//Create service
func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
	}
}

//implement services
type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func (service *userServiceImpl) FindAll() (responses []model.GetProductResponse) {
	users := service.UserRepository.FindAll()
	for _, user := range users {
		responses = append(responses, model.GetProductResponse{
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return responses
}
