package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	ErrSMTPServerNotSet    = errors.New("SMTP server not set")
	ErrSMTPPortNotSet      = errors.New("SMTP port not set")
	ErrSMTPUsernameNotSet  = errors.New("SMTP username not set")
	ErrSMTPPasswordNotSet  = errors.New("SMTP password not set")
	ErrSMTPRecipientNotSet = errors.New("SMTP recipient not set")
)

type SMTPConfig struct {
	Server    string `toml:"smtp_server"`
	Port      int    `toml:"smtp_port"`
	Username  string `toml:"smtp_username"`
	Password  string `toml:"smtp_password"`
	Recipient string `toml:"smtp_recipient"`
}

func mergeSMTPConfigs(config *SMTPConfig) error {
	if server := os.Getenv("REGIOJET_SMTP_SERVER"); server != "" {
		config.Server = server
	}

	if port := os.Getenv("REGIOJET_SMTP_PORT"); port != "" {
		portInt, err := strconv.Atoi(port)
		if err != nil {
			return fmt.Errorf("failed to parse the SMTP port: %w", err)
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

func validateSMTPConfig(config *Config) error {
	if config.SMTP.Server == "" {
		return fmt.Errorf("%w", ErrSMTPServerNotSet)
	}

	if config.SMTP.Port == 0 {
		return fmt.Errorf("%w", ErrSMTPPortNotSet)
	}

	if config.SMTP.Username == "" {
		return fmt.Errorf("%w", ErrSMTPUsernameNotSet)
	}

	if config.SMTP.Password == "" {
		return fmt.Errorf("%w", ErrSMTPPasswordNotSet)
	}

	if config.General.SingleUserMode && config.SMTP.Recipient == "" {
		return fmt.Errorf("%w", ErrSMTPRecipientNotSet)
	}

	return nil
}
