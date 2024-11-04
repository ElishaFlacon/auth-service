package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

const (
	userServicesEnvName = "USER_SERVICES"

	errMsgUserServicesPathNotFound = "User services path not found"
)

type Config interface {
	LoadEnv() error
	GetUserServicesPath() (string, error)
}

type config struct {
	envPath string
}

func NewConfig(envPath string) *config {
	return &config{
		envPath: envPath,
	}
}

func (c *config) LoadEnv() error {
	return godotenv.Load(c.envPath)
}

func (c *config) GetUserServicesPath() (string, error) {
	userServicesPath := os.Getenv(userServicesEnvName)
	if len(userServicesPath) == 0 {
		err := errors.New(errMsgUserServicesPathNotFound)
		return "", err
	}

	return userServicesPath, nil
}
