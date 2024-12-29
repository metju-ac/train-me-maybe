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

	apiConfiguration := openapiclient.NewConfiguration(config.General.ApiBaseUrl)
	apiClient := openapiclient.NewAPIClient(apiConfiguration)

	route, err := handlers.HandleRouteSelection(apiClient)
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
		freeSeats, err := handlers.CheckFreeSeats(apiClient, route.DepartingStation, route.ArrivingStation, route.SelectedRoute, seatClasses)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if freeSeats {
			fmt.Println("Free seats found!")
			notification.EmailNotification(&config.Smtp, route)
			break
		}

		fmt.Printf("No free seats found, checking again in %d seconds...\n", config.General.PollInterval)
		time.Sleep(time.Duration(config.General.PollInterval) * time.Second)
	}
}
