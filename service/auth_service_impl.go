package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/habibiiberahim/go-backend/config"
	"github.com/habibiiberahim/go-backend/entity"
	"github.com/habibiiberahim/go-backend/exception"
	"github.com/habibiiberahim/go-backend/repository"
)

func NewAuthService(userRepository repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		UserRepository: userRepository,
	}
}

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
}

func (service *AuthServiceImpl) Redirect(c *fiber.Ctx) string {
	configuration := config.NewConfiguration()
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

func (service *AuthServiceImpl) Callback(c *fiber.Ctx) string {
	state := c.Query("state")
	code := c.Query("code")
	provider := c.Params("provider")

	// Handle callback and check for errors
	user, _, err := config.Gocial.Handle(state, code)
	exception.PanicIfNeeded(err)

	userData := service.UserRepository.GetOrRegister(provider, user)
	jwtToken := service.CreateToken(&userData)

	// Print in terminal user information
	return jwtToken
}

func (service *AuthServiceImpl) CreateToken(user *entity.User) string {
	configuration := config.NewConfiguration()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().AddDate(0, 0, 7).Unix(),
		"iat":     time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(configuration.Get("APP_SECRET")))
	exception.PanicIfNeeded(err)

	return tokenString
}
