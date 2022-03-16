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
	configuration := config.NewConfiguration()

	//setup fiber
	fiberConfig := config.NewFiberConfig()
	app := fiber.New(fiberConfig)

	//setup database
	database := config.NewGormDatabase(configuration)

	//setup repository
	userRepository := repository.NewUserRepository(database)

	//setup service
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository)

	//setup controller
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)

	//setup routes
	userController.NewRoute(app)
	authController.NewRoute(app)

	//serve port
	err := app.Listen(":8000")
	exception.PanicIfNeeded(err)
}
