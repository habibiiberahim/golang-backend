package controller

import (
	"github.com/gofiber/fiber/v2"
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
	// api := app.Group("/api/v1")
	// api.Get("/auth/:provider", controller.Redirect)
	// api.Get("/auth/:provider/callback", controller.Callback)
}

func (controller *AuthController) Redirect(c fiber.Ctx) error {
	return c.Redirect("")
}

// func (controller *AuthController) Redirect(c *fiber.Ctx) error {
// 	configuration := config.New()
// 	provider := c.Params("provider")

// 	providerSecrets := map[string]map[string]string{
// 		"google": {
// 			"clientID":     configuration.Get("CLIENT_ID_GOOGLE"),
// 			"clientSecret": configuration.Get("CLIENT_SECRET_GOOGLE"),
// 			"redirectURL":  configuration.Get("AUTH_REDIRECT_URL") + "/google/callback",
// 		},
// 	}

// 	providerScopes := map[string][]string{
// 		"google": {},
// 	}

// 	providerData := providerSecrets[provider]
// 	actualScope := providerScopes[provider]
// 	authURL, err := config.Gocial.New().
// 		Driver(provider).
// 		Scopes(actualScope).
// 		Redirect(
// 			providerData["clientID"],
// 			providerData["clientSecret"],
// 			providerData["redirectURL"],
// 		)
// 	exception.PanicIfNeeded(err)

// 	return c.Redirect(authURL)
// }

// func (controller *AuthController) Callback(c *fiber.Ctx) error {
// 	state := c.Query("state")
// 	code := c.Query("code")

// 	// Handle callback and check for errors
// 	user, _, err := config.Gocial.Handle(state, code)
// 	exception.PanicIfNeeded(err)

// 	// Print in terminal user information
// 	return c.JSON(model.WebResponse{
// 		Code:   200,
// 		Status: "Success",
// 		Data:   user.Email,
// 	})

// }
