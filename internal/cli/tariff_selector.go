package cli

import (
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/metju-ac/train-me-maybe/openapi"
)

var ErrTariffNotFound = errors.New("tariff not found")

func formatTariff(tariff *openapi.Tariff) string {
	return fmt.Sprintf("%s (%s)", *tariff.Value, *tariff.Key)
}

func SelectTariff(tariffs []openapi.Tariff) (*openapi.Tariff, error) {
	var selectedTariff string

	formattedTariffs := make([]string, 0, len(tariffs))
	for _, tariff := range tariffs {
		formatted := formatTariff(&tariff)
		formattedTariffs = append(formattedTariffs, formatted)
	}

	selectPrompt := &survey.Select{
		Message: "Select tariff:",
		Options: formattedTariffs,
	}
	err := survey.AskOne(selectPrompt, &selectedTariff)
	if err != nil {
		return nil, fmt.Errorf("failed to select tariff: %w", err)
	}

	for _, tariff := range tariffs {
		formatted := formatTariff(&tariff)
		if formatted == selectedTariff {
			return &tariff, nil
		}
	}

	return nil, fmt.Errorf("%w", ErrTariffNotFound)
}
