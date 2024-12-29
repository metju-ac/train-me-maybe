package config

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/metju-ac/train-me-maybe/internal/lib"
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
		if enabled == "true" {
			config.CreditEnabled = true
		} else if enabled == "false" {
			config.CreditEnabled = false
		} else {
			return errors.New("Invalid value for REGIOJET_AUTH_CREDIT_ENABLED")
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

func validateAuthConfig(config *AuthConfig) error {

	if !config.CreditEnabled {
		return nil
	}

	if config.CreditUser == "" {
		return errors.New("The Credit user must be set")
	}

	if config.CreditPassword == "" {
		return errors.New("The Credit password must be set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	token, err := lib.LoginWithCreditTicket(ctx, config.CreditUser, config.CreditPassword)
	if err != nil {
		return err
	}

	_, err = uuid.Parse(token)
	if err != nil {
		slog.Error("The token is not a valid UUID", "error", err)
		return errors.New("The token is not a valid UUID")
	}

	return nil
}
