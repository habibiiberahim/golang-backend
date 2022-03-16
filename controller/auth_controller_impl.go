package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/go-backend/model"
	"github.com/habibiiberahim/go-backend/service"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) NewRoute(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/auth/:provider", controller.Redirect)
	api.Get("/auth/:provider/callback", controller.Callback)
}

func (controller *AuthControllerImpl) Redirect(c *fiber.Ctx) error {
	authURL := controller.AuthService.Redirect(c)
	return c.Redirect(authURL)
}

func (controller *AuthControllerImpl) Callback(c *fiber.Ctx) error {
	responses := controller.AuthService.Callback(c)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   responses,
	})
}

func (controller *AuthControllerImpl) Login(c *fiber.Ctx) error {
	responses := controller.AuthService.Callback(c)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   responses,
	})
}
