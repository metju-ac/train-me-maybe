package cli

import (
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/metju-ac/train-me-maybe/internal/models"
)

var ErrStationNotFound = errors.New("station not found")

func FormatStation(station *models.StationModel) string {
	return fmt.Sprintf("%s - %s (%s)", station.City, station.StationName, station.Country)
}

func SelectStation(stations []models.StationModel) (*models.StationModel, error) {
	var selectedStation string

	formattedStations := make([]string, 0, len(stations))
	for _, station := range stations {
		formattedStations = append(formattedStations, FormatStation(&station))
	}

	selectPrompt := &survey.Select{
		Message: "Select a station:",
		Options: formattedStations,
	}
	err := survey.AskOne(selectPrompt, &selectedStation)
	if err != nil {
		return nil, fmt.Errorf("failed to select station: %w", err)
	}

	for _, station := range stations {
		formatted := FormatStation(&station)
		if formatted == selectedStation {
			return &station, nil
		}
	}

	return nil, fmt.Errorf("%w", ErrStationNotFound)
}
