package handlers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/metju-ac/train-me-maybe/internal/models"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

var (
	ErrMalformedResponse = errors.New("malformed response from GetRouteFreeSeats")
	ErrNoFreeSeats       = errors.New("no free seats found")
)

func freeSeats(res *openapiclient.RouteSeatsResponse) bool {
	for _, vehicle := range res.Vehicles {
		for _, deck := range vehicle.Decks {
			if len(deck.FreeSeats) > 0 {
				return true
			}
		}
	}
	return false
}

func GetFreeSeatsOnRoute(
	apiClient *openapiclient.APIClient,
	userInput *models.UserInput,
	seatClass string,
) (*openapiclient.RouteSeatsResponse, error) {
	slog.Info("Checking for free seats",
		"departingStation", userInput.DepartingStation.StationID,
		"arrivingStation", userInput.ArrivingStation.StationID,
		"selectedRoutes", userInput.SelectedRouteIDs,
		"seatClasses", userInput.SeatClasses,
	)

	sections := []openapiclient.SimpleSection{*userInput.Section}

	var tariffs []string
	if userInput.TariffKey == "" {
		tariffs = append(tariffs, "REGULAR")
	} else {
		tariffs = append(tariffs, userInput.TariffKey)
	}
	routeSeatsRequest := openapiclient.NewRouteSeatsRequest(sections, tariffs, seatClass)

	route, httpResponse, err := apiClient.RoutesAPI.GetRouteFreeSeats(context.Background()).Request(*routeSeatsRequest).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == http.StatusBadRequest {
			slog.Info("No free seats found (400 response)")
			return nil, ErrNoFreeSeats
		}

		slog.Error("Error calling GetRouteFreeSeats", "error", err)
		return nil, fmt.Errorf("error calling GetRouteFreeSeats: %w", err)
	}
	defer httpResponse.Body.Close()

	if route == nil || len(*route) == 0 {
		return nil, fmt.Errorf("%w", ErrMalformedResponse)
	}

	return &((*route)[0]), nil
}

type CheckFreeSeatsResponse struct {
	SeatClass    string
	HasFreeSeats bool
	SelectedSeat *openapiclient.SelectedSeat
}

func CheckFreeSeats(apiClient *openapiclient.APIClient, userInput *models.UserInput) (*CheckFreeSeatsResponse, error) {
	for _, seatClass := range userInput.SeatClasses {
		route, err := GetFreeSeatsOnRoute(apiClient, userInput, seatClass)
		if err != nil {
			if errors.Is(err, ErrNoFreeSeats) {
				continue
			}
			slog.Error("Error getting free seats", "error", err)
			return nil, fmt.Errorf("error getting free seats: %w", err)
		}

		if route == nil {
			continue
		}

		if freeSeats(route) {
			slog.Info("Free seats found", "seatClass", seatClass)
			return &CheckFreeSeatsResponse{
				SeatClass:    seatClass,
				HasFreeSeats: true,
				SelectedSeat: &route.SelectedSeats[0],
			}, nil
		}
	}

	slog.Info("No free seats found")
	return &CheckFreeSeatsResponse{
		HasFreeSeats: false,
	}, nil
}
