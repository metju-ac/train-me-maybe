package main

import (
	"fmt"
	"os"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/handlers"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func main() {
	fmt.Println("Hello, RegioJet!")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

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
			break
		}

		fmt.Println("No free seats found, checking again in 10 seconds...")
		time.Sleep(10 * time.Second)
	}
}
