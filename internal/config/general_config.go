package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/metju-ac/train-me-maybe/internal/utils"
)

const MinPollInterval = 10

var (
	ErrPollIntervalInvalid          = errors.New("poll interval must be greater than 0")
	ErrPollIntervalTooLow           = errors.New("poll interval is dangerously low, it might get you banned from the API")
	ErrPurchaseCutoffMinutesInvalid = errors.New("purchase cutoff minutes must be greater than 0")
	ErrLowCreditThresholdInvalid    = errors.New("low credit threshold must be non-negative if set")
	ErrAPIBaseURLNotSet             = errors.New("API base URL must be set")
	ErrTicketBeerPriceInvalid       = errors.New("ticket beer price must be greater than 0")
	ErrAllowedOriginsNotSet         = errors.New("allowed origins must be set")
	ErrServerPortInvalid            = errors.New("server port must be greater than 0")
	ErrEncryptionKeyInvalid         = errors.New("encryption key must be 32 bytes")
	ErrLanguagesNotSet              = errors.New("languages must be set")
	ErrInvalidValueForNullableInt   = errors.New("invalid value for NullableInt")
	ErrInvalidTypeForNullableInt    = errors.New("invalid type for NullableInt")
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
			return fmt.Errorf("%w", ErrInvalidValueForNullableInt)
		}
	default:
		return fmt.Errorf("%w", ErrInvalidTypeForNullableInt)
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
	APIBaseURL string `toml:"api_base_url"`

	// Whether to run in single user mode
	SingleUserMode bool `toml:"single_user_mode"`

	// Price of a successful ticket purchase in beers
	TicketBeerPrice float32 `toml:"ticket_beer_price"`

	// Allowed origins for CORS
	AllowedOrigins []string `toml:"allowed_origins"`

	// Port of the server
	ServerPort int `toml:"server_port"`

	// Encryption key for user RJ passwords
	EncryptionKey string `toml:"encryption_key"`

	// Supported languages
	Languages []string `toml:"languages"`
}

func mergeGeneralConfigs(config *GeneralConfig) error {
	if interval := os.Getenv("REGIOJET_POLL_INTERVAL"); interval != "" {
		intervalInt, err := strconv.Atoi(interval)
		if err != nil {
			return fmt.Errorf("failed to convert poll interval to integer: %w", err)
		}
		config.PollInterval = intervalInt
	}

	if cutoff := os.Getenv("REGIOJET_PURCHASE_CUTOFF_MINUTES"); cutoff != "" {
		cutoffInt, err := strconv.Atoi(cutoff)
		if err != nil {
			return fmt.Errorf("failed to convert purchase cutoff minutes to integer: %w", err)
		}
		config.PurchaseCutoffMinutes = cutoffInt
	}

	if threshold := os.Getenv("REGIOJET_LOW_CREDIT_THRESHOLD"); threshold != "" {
		thresholdInt, err := strconv.Atoi(threshold)
		if err != nil {
			return fmt.Errorf("failed to convert low credit threshold to integer: %w", err)
		}
		config.LowCreditThreshold.Value = &thresholdInt
	}

	if url := os.Getenv("REGIOJET_API_BASE_URL"); url != "" {
		config.APIBaseURL = url
	}

	if mode := os.Getenv("REGIOJET_SINGLE_USER_MODE"); mode != "" {
		modeBool, err := strconv.ParseBool(mode)
		if err != nil {
			return fmt.Errorf("failed to convert single user mode to boolean: %w", err)
		}
		config.SingleUserMode = modeBool
	}

	if price := os.Getenv("REGIOJET_TICKET_BEER_PRICE"); price != "" {
		priceFloat, err := strconv.ParseFloat(price, 32)
		if err != nil {
			return fmt.Errorf("failed to convert ticket beer price to float: %w", err)
		}
		config.TicketBeerPrice = float32(priceFloat)
	}

	if origins := os.Getenv("REGIOJET_ALLOWED_ORIGINS"); origins != "" {
		config.AllowedOrigins = append(config.AllowedOrigins, origins)
	}

	if port := os.Getenv("REGIOJET_SERVER_PORT"); port != "" {
		portInt, err := strconv.Atoi(port)
		if err != nil {
			return fmt.Errorf("failed to convert server port to integer: %w", err)
		}
		config.ServerPort = portInt
	}

	if key := os.Getenv("REGIOJET_ENCRYPTION_KEY"); key != "" {
		config.EncryptionKey = key
	}

	if langs := os.Getenv("REGIOJET_LANGUAGES"); langs != "" {
		config.Languages = strings.Split(langs, ",")
	}

	return nil
}

func validateGeneralConfig(config *GeneralConfig) error {
	if config.PollInterval <= 0 {
		return fmt.Errorf("%w", ErrPollIntervalInvalid)
	}

	if config.PollInterval < MinPollInterval {
		return fmt.Errorf("%w", ErrPollIntervalTooLow)
	}

	if config.PurchaseCutoffMinutes <= 0 {
		return fmt.Errorf("%w", ErrPurchaseCutoffMinutesInvalid)
	}

	if config.LowCreditThreshold.Value != nil && *config.LowCreditThreshold.Value < 0 {
		return fmt.Errorf("%w", ErrLowCreditThresholdInvalid)
	}

	if config.APIBaseURL == "" {
		return fmt.Errorf("%w", ErrAPIBaseURLNotSet)
	}

	if config.TicketBeerPrice <= 0 {
		return fmt.Errorf("%w", ErrTicketBeerPriceInvalid)
	}

	if len(config.AllowedOrigins) == 0 {
		return fmt.Errorf("%w", ErrAllowedOriginsNotSet)
	}

	if config.ServerPort <= 0 {
		return fmt.Errorf("%w", ErrServerPortInvalid)
	}

	if len([]byte(config.EncryptionKey)) != utils.KeyLength {
		return fmt.Errorf("%w", ErrEncryptionKeyInvalid)
	}

	if len(config.Languages) == 0 {
		return fmt.Errorf("%w", ErrLanguagesNotSet)
	}

	return nil
}
