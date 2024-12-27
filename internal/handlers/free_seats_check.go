package handlers

import (
	"context"
	"fmt"
	"log/slog"

	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func freeSeats(res []openapiclient.RouteSeatsResponse) bool {
	for _, routeSeatsResponse := range res {
		for _, vehicle := range routeSeatsResponse.Vehicles {
			for _, deck := range vehicle.Decks {
				if len(deck.FreeSeats) > 0 {
					return true
				}
			}
		}
	}
	return false
}

func CheckFreeSeats(apiClient *openapiclient.APIClient, departingStation, arrivingStation int64, selectedRoutes []int64, seatClasses []string) (bool, error) {
	slog.Info("Checking for free seats", "departingStation", departingStation, "arrivingStation", arrivingStation, "selectedRoutes", selectedRoutes, "seatClasses", seatClasses)
	sections := []openapiclient.SimpleSection{
		*openapiclient.NewSimpleSection(selectedRoutes[0], departingStation, arrivingStation),
	}
	tariffs := []string{"REGULAR"}

	for _, seatClass := range seatClasses {
		routeSeatsRequest := openapiclient.NewRouteSeatsRequest(sections, tariffs, seatClass)

		route, _, err := apiClient.RoutesAPI.GetRouteFreeSeats(context.Background()).Request(*routeSeatsRequest).Execute()
		if err != nil {
			slog.Error("Error calling GetRouteFreeSeats", "error", err)
			return false, fmt.Errorf("error calling GetRouteFreeSeats: %v", err)
		}

		if freeSeats(*route) {
			slog.Info("Free seats found", "seatClass", seatClass)
			return true, nil
		}
	}

	slog.Info("No free seats found")
	return false, nil
}
