package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	SMTP    SMTPConfig    `toml:"smtp"`
	General GeneralConfig `toml:"general"`
	Auth    AuthConfig    `toml:"auth"`
	DB      DBConfig      `toml:"db"`
}

func LoadConfig() (*Config, error) {
	slog.Info("Loading configuration")

	// first, try to load TOML config file
	var config Config

	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		slog.Error("Failed to load TOML config file", "error", err)
		return nil, fmt.Errorf("failed to load TOML config file: %w", err)
	}

	// Check if .env file exists and load it if present
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			slog.Error("Failed to load .env file", "error", err)
			return nil, fmt.Errorf("failed to load .env file: %w", err)
		}
		slog.Info("Loaded .env file")
	} else {
		slog.Info(".env file not found, using environment variables from the system")
	}

	// then, if something is in the environment variables, overwrite the config
	if err := mergeSMTPConfigs(&config.SMTP); err != nil {
		slog.Error("Failed to merge SMTP configs", "error", err)
		return nil, err
	}

	if err := mergeGeneralConfigs(&config.General); err != nil {
		slog.Error("Failed to merge general configs", "error", err)
		return nil, err
	}

	if err := mergeAuthConfigs(&config.Auth); err != nil {
		slog.Error("Failed to merge auth configs", "error", err)
		return nil, err
	}

	if err := mergeDBConfigs(&config.DB); err != nil {
		slog.Error("Failed to merge db configs", "error", err)
		return nil, err
	}

	// more sections will go here

	// validate the configs
	if err := validateSMTPConfig(&config); err != nil {
		slog.Error("Failed to validate SMTP config", "error", err)
		return nil, err
	}

	if err := validateGeneralConfig(&config.General); err != nil {
		slog.Error("Failed to validate general config", "error", err)
		return nil, err
	}

	if err := validateAuthConfig(&config); err != nil {
		slog.Error("Failed to validate auth config", "error", err)
		return nil, err
	}

	if err := validateDBConfig(&config); err != nil {
		slog.Error("Failed to validate db config", "error", err)
		return nil, err
	}

	// more validations will go here

	slog.Info("Configuration loaded successfully")
	return &config, nil
}
