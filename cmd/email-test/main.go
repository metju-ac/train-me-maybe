package main

import (
	"fmt"
	"os"

	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/handlers"
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

	notification.EmailNotification(&config.Smtp, &handlers.HandleRouteSelectionResponse{
		DepartingStation: test_station,
		ArrivingStation:  test_station,
		SelectedRoute:    openapi.NewSimpleRouteWithDefaults(),
	})
}
