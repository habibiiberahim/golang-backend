package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiberahim/go-backend/config"
	"github.com/habibiiberahim/go-backend/exception"
)

func NewAuthService() AuthService {
	return &authServiceImpl{}
}

type authServiceImpl struct {
	AuthService
}

func (service *authServiceImpl) Redirect(c *fiber.Ctx) string {
	configuration := config.New()
	provider := c.Params("provider")

	providerSecrets := map[string]map[string]string{
		"google": {
			"clientID":     configuration.Get("CLIENT_ID_GOOGLE"),
			"clientSecret": configuration.Get("CLIENT_SECRET_GOOGLE"),
			"redirectURL":  configuration.Get("AUTH_REDIRECT_URL") + "/google/callback",
		},
	}

	providerScopes := map[string][]string{
		"google": {},
	}

	providerData := providerSecrets[provider]
	actualScope := providerScopes[provider]
	authURL, err := config.Gocial.New().
		Driver(provider).
		Scopes(actualScope).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)
	exception.PanicIfNeeded(err)

	return authURL
}

func (service *authServiceImpl) Callback(c *fiber.Ctx) string {
	state := c.Query("state")
	code := c.Query("code")

	// Handle callback and check for errors
	user, _, err := config.Gocial.Handle(state, code)
	exception.PanicIfNeeded(err)

	// Print in terminal user information
	return user.FullName
}
