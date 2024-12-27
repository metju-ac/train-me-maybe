package config

import (
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	Smtp SmtpConfig
}

func LoadConfig() (*Config, error) {
	// first, try to load TOML config file
	var config Config

	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		return nil, err
	}

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// then, if something is in the environment variables, overwrite the config
	if err := mergeSmtpConfigs(&config.Smtp); err != nil {
		return nil, err
	}

	// more sections will go here

	// validate the configs
	if err := validateSmtpConfig(&config.Smtp); err != nil {
		return nil, err
	}

	// more validations will go here

	return &config, nil
}
