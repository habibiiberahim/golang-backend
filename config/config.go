package config

import (
	"os"

	"github.com/habibiiberahim/go-backend/exception"
	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func NewConfiguration() Config {
	err := godotenv.Load()

	exception.PanicIfNeeded(err)
	return &configImpl{}
}
