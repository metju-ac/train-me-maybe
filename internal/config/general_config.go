package config

import (
	"errors"
	"os"
	"strconv"
)

type GeneralConfig struct {
	// Polling interval of the regiojet API in seconds
	PollInterval int `toml:"poll_interval"`

	// Base URL of the regiojet API. E.g. 'https://brn-ybus-pubapi.sa.cz/restapi'
	ApiBaseUrl string `toml:"api_base_url"`
}

func mergeGeneralConfigs(config *GeneralConfig) error {

	if interval := os.Getenv("REGIOJET_POLL_INTERVAL"); interval != "" {
		intervalInt, err := strconv.Atoi(interval)
		if err != nil {
			return err
		}
		config.PollInterval = intervalInt
	}

	if url := os.Getenv("REGIOJET_API_BASE_URL"); url != "" {
		config.ApiBaseUrl = url
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

	if config.ApiBaseUrl == "" {
		return errors.New("API base URL must be set")
	}

	return nil
}
