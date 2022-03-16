package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/go-backend/middleware"
	"github.com/habibiiberahim/go-backend/model"
	"github.com/habibiiberahim/go-backend/service"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) NewRoute(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/users", middleware.Protected(), controller.FindAll)
}

func (controller *UserControllerImpl) FindAll(c *fiber.Ctx) error {
	// responses := controller.UserService.FindAll()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   "responses",
	})
}
