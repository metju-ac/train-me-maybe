package config

import (
	"errors"
	"os"
	"strconv"
)

type NullableInt struct {
	Value *int
}

func (ni *NullableInt) UnmarshalTOML(data interface{}) error {
	if data == nil {
		ni.Value = nil
		return nil
	}

	switch v := data.(type) {
	case int64:
		val := int(v)
		ni.Value = &val
	case string:
		if v == "" {
			ni.Value = nil
		} else {
			return errors.New("invalid value for NullableInt")
		}
	default:
		return errors.New("invalid type for NullableInt")
	}

	return nil
}

type GeneralConfig struct {
	// Polling interval of the regiojet API in seconds
	PollInterval int `toml:"poll_interval"`

	// Minutes before departure to allow purchase
	PurchaseCutoffMinutes int `toml:"purchase_cutoff_minutes"`

	// Threshold for low credit warning
	LowCreditThreshold NullableInt `toml:"low_credit_threshold"`

	// Base URL of the regiojet API. E.g. 'https://brn-ybus-pubapi.sa.cz/restapi'
	ApiBaseUrl string `toml:"api_base_url"`

	// Whether to run in single user mode
	SingleUserMode bool `toml:"single_user_mode"`
}

func mergeGeneralConfigs(config *GeneralConfig) error {

	if interval := os.Getenv("REGIOJET_POLL_INTERVAL"); interval != "" {
		intervalInt, err := strconv.Atoi(interval)
		if err != nil {
			return err
		}
		config.PollInterval = intervalInt
	}

	if cutoff := os.Getenv("REGIOJET_PURCHASE_CUTOFF_MINUTES"); cutoff != "" {
		cutoffInt, err := strconv.Atoi(cutoff)
		if err != nil {
			return err
		}
		config.PurchaseCutoffMinutes = cutoffInt
	}

	if threshold := os.Getenv("REGIOJET_LOW_CREDIT_THRESHOLD"); threshold != "" {
		thresholdInt, err := strconv.Atoi(threshold)
		if err != nil {
			return err
		}
		config.LowCreditThreshold.Value = &thresholdInt
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

	if config.PurchaseCutoffMinutes <= 0 {
		return errors.New("Purchase cutoff minutes must be greater than 0.")
	}

	if config.LowCreditThreshold.Value != nil && *config.LowCreditThreshold.Value < 0 {
		return errors.New("Low credit threshold must be non-negative if set")
	}

	if config.ApiBaseUrl == "" {
		return errors.New("API base URL must be set")
	}

	return nil
}
