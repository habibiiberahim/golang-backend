package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/go-backend/config"
	"github.com/habibiiberahim/go-backend/controller"
	"github.com/habibiiberahim/go-backend/exception"
	"github.com/habibiiberahim/go-backend/repository"
	"github.com/habibiiberahim/go-backend/service"
)

func main() {
	//setup config
	configuration := config.New()

	//setup database
	database := config.NewGormDatabase(configuration)

	//setup fiber
	app := fiber.New(config.NewFiberConfig())

	//setup repository
	userRepository := repository.NewUserRepository(database)

	//setup service
	userService := service.NewUserService(&userRepository)
	authService := service.NewAuthService(&userRepository)

	//setup controller
	userController := controller.NewUserController(&userService)
	authController := controller.NewAuthController(&authService)
	//setup routes
	userController.Route(app)
	authController.Route(app)

	err := app.Listen(":3001")

	exception.PanicIfNeeded(err)
}
