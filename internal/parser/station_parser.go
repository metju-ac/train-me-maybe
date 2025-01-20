package parser

import (
	"log/slog"

	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/openapi"
)

const stationsMaximum = 300

func TransformStations(countries []openapi.Country) []models.StationModel {
	result := make([]models.StationModel, 0, stationsMaximum)

	for _, country := range countries {
		for _, city := range country.Cities {
			for _, station := range city.Stations {
				isTrainStation := false
				isBusStation := false
				for _, stationType := range station.StationsTypes {
					if stationType == "TRAIN_STATION" {
						isTrainStation = true
					}
					if stationType == "BUS_STATION" {
						isBusStation = true
					}
				}

				if !isTrainStation && !isBusStation {
					slog.Warn("Station is neither a train station nor a bus station", "stationName", station.Name, "stationID", station.Id)
				}

				result = append(result, models.StationModel{
					Country:        country.Country,
					City:           city.Name,
					StationID:      station.Id,
					StationName:    station.Name,
					IsTrainStation: isTrainStation,
					IsBusStation:   isBusStation,
				})
			}
		}
	}

	slog.Info("Transformed stations", "count", len(result))
	return result
}
