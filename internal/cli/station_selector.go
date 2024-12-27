package cli

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/metju-ac/train-me-maybe/internal/models"
)

func formatStation(station models.StationModel) string {
	return fmt.Sprintf("%s - %s (%s)", station.City, station.StationName, station.Country)
}

func SelectStation(stations []models.StationModel) (int64, error) {
	var selectedStation string

	formattedStations := make([]string, 0, len(stations))
	for _, station := range stations {
		formattedStations = append(formattedStations, formatStation(station))
	}

	selectPrompt := &survey.Select{
		Message: "Select a station:",
		Options: formattedStations,
	}
	err := survey.AskOne(selectPrompt, &selectedStation)
	if err != nil {
		return 0, err
	}

	for _, station := range stations {
		formatted := formatStation(station)
		if formatted == selectedStation {
			return station.StationID, nil
		}
	}

	return 0, fmt.Errorf("station not found")
}
