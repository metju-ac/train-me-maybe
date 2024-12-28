package config

import (
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"log/slog"
)

type Config struct {
	Smtp SmtpConfig
}

func LoadConfig() (*Config, error) {
	slog.Info("Loading configuration")

	// first, try to load TOML config file
	var config Config

	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		slog.Error("Failed to load TOML config file", "error", err)
		return nil, err
	}

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		slog.Error("Failed to load .env file", "error", err)
		return nil, err
	}

	// then, if something is in the environment variables, overwrite the config
	if err := mergeSmtpConfigs(&config.Smtp); err != nil {
		slog.Error("Failed to merge SMTP configs", "error", err)
		return nil, err
	}

	// more sections will go here

	// validate the configs
	if err := validateSmtpConfig(&config.Smtp); err != nil {
		slog.Error("Failed to validate SMTP config", "error", err)
		return nil, err
	}

	// more validations will go here

	slog.Info("Configuration loaded successfully")
	return &config, nil
}
