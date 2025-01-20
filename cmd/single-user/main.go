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
	configuration, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	apiConfiguration := openapiclient.NewConfiguration(configuration.General.APIBaseURL)
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
	if configuration.Auth.CreditEnabled {
		tariff, err = handlers.HandleTariffSelection(apiClient)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}

	userInput := &models.UserInput{
		DepartingStation:   route.DepartingStation,
		ArrivingStation:    route.ArrivingStation,
		SelectedRouteIDs:   route.SelectedRoute.Id,
		SeatClasses:        seatClasses,
		TariffKey:          *tariff.Key,
		Section:            nil,
		RouteDetail:        nil,
		CreditUserNumber:   configuration.Auth.CreditUser,
		CreditUserPassword: configuration.Auth.CreditPassword,
		UserEmail:          configuration.SMTP.Recipient,
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

	purchaseCutoffDuration := time.Duration(configuration.General.PurchaseCutoffMinutes) * time.Minute

	for {
		if time.Now().After(userInput.RouteDetail.DepartureTime) {
			fmt.Println("The train has already departed.")
			break
		}

		freeSeatsResp, err := handlers.CheckFreeSeats(apiClient, userInput)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if freeSeatsResp == nil || !freeSeatsResp.HasFreeSeats {
			fmt.Printf("No free seats found, checking again in %d seconds...\n", configuration.General.PollInterval)
			time.Sleep(time.Duration(configuration.General.PollInterval) * time.Second)
			continue
		}

		fmt.Println("Free seats found!")

		if configuration.Auth.CreditEnabled && time.Now().Add(purchaseCutoffDuration).Before(userInput.RouteDetail.DepartureTime) {
			response, ticket, err := purchase.AutoPurchaseTicket(apiClient, configuration, userInput, freeSeatsResp)
			if err != nil {
				fmt.Println("Error purchasing ticket:", err)
				os.Exit(1)
			}

			notification.EmailNotificationTicketBought(&configuration.SMTP, userInput, ticket)

			if configuration.General.LowCreditThreshold.Value != nil && *configuration.General.LowCreditThreshold.Value >= int(response.Amount) {
				notification.EmailNotificationLowCredit(&configuration.SMTP, userInput, response.Amount, response.Currency)
			}

			break
		}

		notification.EmailNotificationFreeSeats(&configuration.SMTP, userInput)
		break
	}
}
