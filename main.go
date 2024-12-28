package main

import (
	"fmt"
	"os"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/handlers"
	"github.com/metju-ac/train-me-maybe/internal/notification"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	apiConfiguration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(apiConfiguration)

	departingStation, arrivingStation, selectedRoutes, err := handlers.HandleRouteSelection(apiClient)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	seatClasses, err := handlers.HandleSeatClassSelection(apiClient)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for {
		freeSeats, err := handlers.CheckFreeSeats(apiClient, departingStation, arrivingStation, selectedRoutes, seatClasses)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if freeSeats {
			fmt.Println("Free seats found!")
			notification.EmailNotification(&config.Smtp, departingStation, arrivingStation, selectedRoutes, seatClasses)
			break
		}

		fmt.Println("No free seats found, checking again in 10 seconds...")
		time.Sleep(time.Duration(config.General.PollInterval) * time.Second)
	}
}
