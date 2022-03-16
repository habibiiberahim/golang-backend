//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/habibiiberahim/go-backend/config"
	"github.com/habibiiberahim/go-backend/controller"
	"github.com/habibiiberahim/go-backend/repository"
	"github.com/habibiiberahim/go-backend/service"
)

func InitializedServer() controller.AuthController {
	wire.Build(
		//config inject
		config.NewConfiguration,
		config.NewGormDatabase,

		//repo inject
		repository.NewUserRepository,
		wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),

		service.NewUserService,
		wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
		service.NewAuthService,
		wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),

		controller.NewUserController,
		wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
		controller.NewAuthController,
		wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
	)
	return nil
}
