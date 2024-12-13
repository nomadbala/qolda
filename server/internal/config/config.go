package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
)

var (
	ErrFailedGetWorkingDirectory = errors.New("unable to get working directory")
	ErrFailedLoadENV             = errors.New("failed to load environment variables")
	ErrFailedLoadResendConfig    = errors.New("failed to load resend configs")
)

type Config struct {
	Resend ResendConfig
}

type ResendConfig struct {
	APIKey string
}

func New() (Config, error) {
	config := Config{}

	root, err := os.Getwd()
	if err != nil {
		return config, ErrFailedGetWorkingDirectory
	}

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		return config, ErrFailedLoadENV
	}

	if err = envconfig.Process("RESEND", &config.Resend); err != nil {
		return config, ErrFailedLoadResendConfig
	}

	return config, nil
}
