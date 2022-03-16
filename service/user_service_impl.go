package service

import (
	"github.com/habibiiberahim/go-backend/repository"
)

//Create service
func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

//implement services
type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

// func (service *UserServiceImpl) FindAll() (responses []model.GetUserResponse) {
// 	users := service.UserRepository.FindAll()
// 	for _, user := range users {
// 		responses = append(responses, model.GetUserResponse{
// 			FullName: user.FullName,
// 			Email:    user.Email,
// 		})
// 	}
// 	return responses
// }
