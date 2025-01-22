package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/internal/notification"
	"github.com/metju-ac/train-me-maybe/openapi"
)

func main() {
	configuration, err := config.LoadConfig()
	if err != nil {
		slog.Error("Error loading configuration", "error", err)
		os.Exit(1)
	}

	testStation := &models.StationModel{
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

	notification.EmailNotificationFreeSeats(&configuration.SMTP, &models.UserInput{
		DepartingStation: testStation,
		ArrivingStation:  testStation,
		SelectedRouteIDs: "123",
		SeatClasses:      []string{"SEAT_CLASS"},
		TariffKey:        "TARIFF_KEY",
		Section:          nil,
		RouteDetail:      &testRoute,
		UserEmail:        configuration.SMTP.Recipient,
	})
}
