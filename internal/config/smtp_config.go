package config

import (
	"errors"
	"os"
	"strconv"
)

type SmtpConfig struct {
	Server    string
	Port      int
	Username  string
	Password  string
	Recipient string
}

func mergeSmtpConfigs(config *SmtpConfig) error {
	if server := os.Getenv("REGIOJET_SMTP_SERVER"); server != "" {
		config.Server = server
	}

	if port := os.Getenv("REGIOJET_SMTP_PORT"); port != "" {
		portInt, err := strconv.Atoi(port)
		if err != nil {
			return err
		}
		config.Port = portInt
	}

	if username := os.Getenv("REGIOJET_SMTP_USERNAME"); username != "" {
		config.Username = username
	}

	if password := os.Getenv("REGIOJET_SMTP_PASSWORD"); password != "" {
		config.Password = password
	}

	if recipient := os.Getenv("REGIOJET_SMTP_RECIPIENT"); recipient != "" {
		config.Recipient = recipient
	}

	return nil
}

func validateSmtpConfig(config *Config) error {
	if config.Smtp.Server == "" {
		return errors.New("SMTP server not set")
	}

	if config.Smtp.Port == 0 {
		return errors.New("SMTP port not set")
	}

	if config.Smtp.Username == "" {
		return errors.New("SMTP username not set")
	}

	if config.Smtp.Password == "" {
		return errors.New("SMTP password not set")
	}

	if config.General.SingleUserMode && config.Smtp.Recipient == "" {
		return errors.New("SMTP recipient not set")
	}

	return nil
}
