package main

import (
	"fmt"
	"os"

	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/internal/notification"
	"github.com/metju-ac/train-me-maybe/openapi"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	test_station := &models.StationModel{
		Country:        "COUNTRY",
		City:           "CITY",
		StationID:      1,
		StationName:    "STATION_NAME",
		IsTrainStation: true,
		IsBusStation:   false,
	}

	notification.EmailNotificationFreeSeats(&config.Smtp, &models.UserInput{
		DepartingStation: test_station,
		ArrivingStation:  test_station,
		SelectedRoute:    openapi.NewSimpleRouteWithDefaults(),
		SeatClasses:      []string{"SEAT_CLASS"},
		Tariff:           nil,
		Section:          nil,
	})
}
