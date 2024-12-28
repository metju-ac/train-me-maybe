package handlers

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/metju-ac/train-me-maybe/internal/cli"
	"github.com/metju-ac/train-me-maybe/internal/models"
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

func CheckFreeSeats(apiClient *openapiclient.APIClient, departingStation, arrivingStation *models.StationModel, selectedRoute *openapiclient.SimpleRoute, seatClasses []string) (bool, error) {
	slog.Info("Checking for free seats", "departingStation", departingStation.StationID, "arrivingStation", arrivingStation.StationID, "selectedRoutes", selectedRoute.Id, "seatClasses", seatClasses)

	sectionIds, err := cli.GetSectionIdsFromRoute(selectedRoute)
	if err != nil {
		slog.Error("Error getting section IDs", "error", err)
		return false, fmt.Errorf("error getting section IDs: %v", err)
	}

	// For now, we only support exactly one section of the given route
	if len(sectionIds) != 1 {
		slog.Error("Expected one section ID, got", "sectionIds", sectionIds)
		return false, fmt.Errorf("expected one section ID, got %v", sectionIds)
	}

	sections := []openapiclient.SimpleSection{
		*openapiclient.NewSimpleSection(sectionIds[0], departingStation.StationID, arrivingStation.StationID),
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
