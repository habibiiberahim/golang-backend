package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/go-backend/model"
	"github.com/habibiiberahim/go-backend/service"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{
		UserService: *userService,
	}
}

func (controller *UserController) Route(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/users", controller.FindAll)
}

func (controller *UserController) FindAll(c *fiber.Ctx) error {
	responses := controller.UserService.FindAll()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   responses,
	})
}
