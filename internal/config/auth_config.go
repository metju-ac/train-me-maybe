package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/metju-ac/train-me-maybe/internal/lib"
)

var (
	ErrInvalidCreditEnabledValue = errors.New("invalid value for REGIOJET_AUTH_CREDIT_ENABLED")
	ErrCreditUserNotSet          = errors.New("the Credit user must be set")
	ErrCreditPasswordNotSet      = errors.New("the Credit password must be set")
	ErrInvalidUUIDToken          = errors.New("the token is not a valid UUID")
)

type AuthConfig struct {
	// whether or not we should automatically purchase the tickets
	CreditEnabled bool `toml:"credit_enabled"`
	// The number of the "kreditova jizdenka"
	CreditUser string `toml:"credit_user"`
	// The password of the "kreditova jizdenka"
	CreditPassword string `toml:"credit_password"`
}

func mergeAuthConfigs(config *AuthConfig) error {
	if enabled := os.Getenv("REGIOJET_AUTH_CREDIT_ENABLED"); enabled != "" {
		switch enabled {
		case "true":
			config.CreditEnabled = true
		case "false":
			config.CreditEnabled = false
		default:
			return fmt.Errorf("%w", ErrInvalidCreditEnabledValue)
		}
	}

	if user := os.Getenv("REGIOJET_AUTH_CREDIT_USER"); user != "" {
		config.CreditUser = user
	}

	if password := os.Getenv("REGIOJET_AUTH_CREDIT_PASSWORD"); password != "" {
		config.CreditPassword = password
	}

	return nil
}

func validateAuthConfig(config *Config) error {
	if !config.Auth.CreditEnabled || !config.General.SingleUserMode {
		return nil
	}

	if config.Auth.CreditUser == "" {
		return fmt.Errorf("%w", ErrCreditUserNotSet)
	}

	if config.Auth.CreditPassword == "" {
		return fmt.Errorf("%w", ErrCreditPasswordNotSet)
	}

	token, err := lib.LoginWithCreditTicket(config.General.APIBaseURL, config.Auth.CreditUser, config.Auth.CreditPassword)
	if err != nil {
		return fmt.Errorf("failed to login with credit ticket: %w", err)
	}

	_, err = uuid.Parse(token)
	if err != nil {
		slog.Error("The token is not a valid UUID", "error", err)
		return fmt.Errorf("%w", ErrInvalidUUIDToken)
	}

	return nil
}
