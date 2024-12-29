package main

import (
	"fmt"
	"os"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/handlers"
	"github.com/metju-ac/train-me-maybe/internal/lib"
	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/internal/notification"
	"github.com/metju-ac/train-me-maybe/internal/purchase"
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

	tariff := handlers.DefaultTariff()
	if config.Auth.CreditEnabled {
		tariff, err = handlers.HandleTariffSelection(apiClient)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}

	userInput := &models.UserInput{
		DepartingStation: route.DepartingStation,
		ArrivingStation:  route.ArrivingStation,
		SelectedRoute:    route.SelectedRoute,
		SeatClasses:      seatClasses,
		Tariff:           tariff,
		Section:          nil,
		RouteDetail:      nil,
	}

	routeDetail, err := handlers.FetchRouteDetail(apiClient, userInput)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	userInput.RouteDetail = routeDetail

	section, err := lib.GetSection(userInput)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	userInput.Section = section

	for {
		freeSeatsResp, err := handlers.CheckFreeSeats(apiClient, userInput)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if freeSeatsResp == nil || !freeSeatsResp.HasFreeSeats {
			fmt.Printf("No free seats found, checking again in %d seconds...\n", config.General.PollInterval)
			time.Sleep(time.Duration(config.General.PollInterval) * time.Second)
			continue
		}

		fmt.Println("Free seats found!")
		purchase.AutoPurchaseTicket(apiClient, config, userInput, freeSeatsResp)
		notification.EmailNotification(&config.Smtp, userInput)
		break
	}
}
