package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/go-backend/model"
	"github.com/habibiiberahim/go-backend/service"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService *service.AuthService) AuthController {
	return AuthController{
		AuthService: *authService,
	}
}

func (controller *AuthController) Route(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/auth/:provider", controller.Redirect)
	api.Get("/auth/:provider/callback", controller.Callback)
}

func (controller *AuthController) Redirect(c *fiber.Ctx) error {
	authURL := controller.AuthService.Redirect(c)
	return c.Redirect(authURL)
}

func (controller *AuthController) Callback(c *fiber.Ctx) error {
	responses := controller.AuthService.Callback(c)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   responses,
	})
}
