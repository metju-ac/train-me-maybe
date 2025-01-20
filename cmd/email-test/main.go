package main

import (
	"fmt"
	"os"
	"time"

	"github.com/metju-ac/train-me-maybe/openapi"

	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/internal/notification"
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

	testRoute := openapi.Route{
		DepartureTime: time.Date(2023, time.November, 10, 15, 0, 0, 0, time.UTC),
	}

	notification.EmailNotificationFreeSeats(&config.Smtp, &models.UserInput{
		DepartingStation: test_station,
		ArrivingStation:  test_station,
		SelectedRouteIds: "123",
		SeatClasses:      []string{"SEAT_CLASS"},
		TariffKey:        "TARIFF_KEY",
		Section:          nil,
		RouteDetail:      &testRoute,
		UserEmail:        config.Smtp.Recipient,
	})
}
