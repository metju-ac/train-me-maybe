package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	ErrDBHostNotSet     = errors.New("the database host must be set")
	ErrDBPortNotSet     = errors.New("the database port must be set")
	ErrDBNameNotSet     = errors.New("the database name must be set")
	ErrDBUserNotSet     = errors.New("the database user must be set")
	ErrDBPasswordNotSet = errors.New("the database password must be set")
)

type DBConfig struct {
	// The database host
	Host string `toml:"db_host"`
	// The database port
	Port int `toml:"db_port"`
	// The database name
	Name string `toml:"db_name"`
	// The database user
	User string `toml:"db_user"`
	// The database password
	Password string `toml:"db_password"`
}

func mergeDBConfigs(config *DBConfig) error {
	if host := os.Getenv("REGIOJET_DB_HOST"); host != "" {
		config.Host = host
	}

	if port := os.Getenv("REGIOJET_DB_PORT"); port != "" {
		portInt, err := strconv.Atoi(port)
		if err != nil {
			return fmt.Errorf("failed to parse the database port: %w", err)
		}
		config.Port = portInt
	}

	if name := os.Getenv("REGIOJET_DB_NAME"); name != "" {
		config.Name = name
	}

	if user := os.Getenv("REGIOJET_DB_USER"); user != "" {
		config.User = user
	}

	if password := os.Getenv("REGIOJET_DB_PASSWORD"); password != "" {
		config.Password = password
	}

	return nil
}

func validateDBConfig(config *Config) error {
	if config.DB.Host == "" {
		return fmt.Errorf("%w", ErrDBHostNotSet)
	}

	if config.DB.Port == 0 {
		return fmt.Errorf("%w", ErrDBPortNotSet)
	}

	if config.DB.Name == "" {
		return fmt.Errorf("%w", ErrDBNameNotSet)
	}

	if config.DB.User == "" {
		return fmt.Errorf("%w", ErrDBUserNotSet)
	}

	if config.DB.Password == "" {
		return fmt.Errorf("%w", ErrDBPasswordNotSet)
	}

	return nil
}
