package config

import (
	"errors"
	"os"
	"strconv"
)

type GeneralConfig struct {
	// Polling interval of the regiojet API in seconds
	PollInterval int `toml:"poll_interval"`
}

func mergeGeneralConfigs(config *GeneralConfig) error {

	if interval := os.Getenv("REGIOJET_POLL_INTERVAL"); interval != "" {
		intervalInt, err := strconv.Atoi(interval)
		if err != nil {
			return err
		}
		config.PollInterval = intervalInt
	}

	return nil
}

func validateGeneralConfig(config *GeneralConfig) error {
	if config.PollInterval <= 0 {
		return errors.New("Poll interval must be greater than 0")
	}

	if config.PollInterval < 10 {
		return errors.New("Poll interval is dangerously low, it might get you banned from the API")
	}

	return nil
}
