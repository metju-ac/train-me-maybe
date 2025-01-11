package config

import (
	"errors"
	"os"
	"strconv"
)

type DbConfig struct {
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

func mergeDbConfigs(config *DbConfig) error {
	if host := os.Getenv("REGIOJET_DB_HOST"); host != "" {
		config.Host = host
	}

	if port := os.Getenv("REGIOJET_DB_PORT"); port != "" {
		if portInt, err := strconv.Atoi(port); err != nil {
			return err
		} else {
			config.Port = portInt
		}
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

func validateDbConfig(config *Config) error {
	if config.Db.Host == "" {
		return errors.New("The database host must be set")
	}

	if config.Db.Port == 0 {
		return errors.New("The database port must be set")
	}

	if config.Db.Name == "" {
		return errors.New("The database name must be set")
	}

	if config.Db.User == "" {
		return errors.New("The database user must be set")
	}

	if config.Db.Password == "" {
		return errors.New("The database password must be set")
	}

	return nil
}
