package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/cli"
	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/internal/parser"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func FetchStations(ctx context.Context, apiClient *openapiclient.APIClient, lang string) ([]models.StationModel, error) {
	slog.Info("Fetching locations")
	data, httpRes, err := apiClient.ConstsAPI.GetLocations(ctx).XLang(lang).Execute()
	if err != nil {
		slog.Error("Failed to fetch station data", "error", err)
		return nil, fmt.Errorf("failed to fetch station data: %w", err)
	}

	slog.Info("Successfully fetched locations", "statusCode", httpRes.StatusCode)
	return parser.TransformStations(data), nil
}

func selectDepartingStation(parsedStations []models.StationModel) (*models.StationModel, error) {
	slog.Info("Selecting departing station")
	fmt.Println("Select the departing station:")

	departingStation, err := cli.SelectStation(parsedStations)
	if err != nil {
		slog.Error("Failed to select departing station", "error", err)
		return nil, fmt.Errorf("failed to select departing station: %w", err)
	}

	slog.Info("Successfully selected departing station", "stationID", departingStation.StationID)
	fmt.Printf("You selected departing station: %+v\n", departingStation.StationID)
	return departingStation, nil
}

func selectArrivingStation(parsedStations []models.StationModel) (*models.StationModel, error) {
	slog.Info("Selecting arriving station")
	fmt.Println("Select the arriving station:")

	arrivingStation, err := cli.SelectStation(parsedStations)
	if err != nil {
		slog.Error("Failed to select arriving station", "error", err)
		return nil, fmt.Errorf("failed to select arriving station: %w", err)
	}

	slog.Info("Successfully selected arriving station", "stationID", arrivingStation.StationID)
	fmt.Printf("You selected arriving station: %+v\n", arrivingStation.StationID)
	return arrivingStation, nil
}

func selectDepartureDate() (string, error) {
	slog.Info("Selecting departure date")
	fmt.Println("Select the departure date:")

	departureDate, err := cli.SelectDate()
	if err != nil {
		slog.Error("Failed to select departure date", "error", err)
		return "", fmt.Errorf("failed to select departure date: %w", err)
	}

	slog.Info("Successfully selected departure date", "date", departureDate)
	fmt.Printf("You selected departure date: %s\n", departureDate)
	return departureDate, nil
}

func FetchRoutes(apiClient *openapiclient.APIClient, departingStation, arrivingStation int64, departureDate string) ([]openapiclient.SimpleRoute, error) {
	slog.Info("Fetching routes", "departingStation", departingStation, "arrivingStation", arrivingStation, "departureDate", departureDate)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	routes, httpRes, err := apiClient.RoutesAPI.SimpleSearchRoutes(ctx).
		FromLocationId(departingStation).
		FromLocationType("STATION").
		ToLocationId(arrivingStation).
		ToLocationType("STATION").
		DepartureDate(departureDate).
		Execute()
	if err != nil {
		slog.Error("Failed to fetch routes", "error", err)
		return nil, fmt.Errorf("failed to fetch routes: %w", err)
	}

	slog.Info("Successfully fetched routes", "statusCode", httpRes.StatusCode)
	return routes.Routes, nil
}

func selectRoute(routes []openapiclient.SimpleRoute) (*openapiclient.SimpleRoute, error) {
	slog.Info("Selecting route")
	fmt.Println("Select the route:")

	selectedRoute, err := cli.SelectRoute(routes)
	if err != nil {
		slog.Error("Failed to select route", "error", err)
		return nil, fmt.Errorf("failed to select route: %w", err)
	}

	slog.Info("Successfully selected route", "routeIDs", selectedRoute.Id)
	fmt.Printf("You selected routes: %+v\n", selectedRoute.Id)
	return selectedRoute, nil
}

type HandleRouteSelectionResponse struct {
	DepartingStation *models.StationModel
	ArrivingStation  *models.StationModel
	SelectedRoute    *openapiclient.SimpleRoute
}

func HandleRouteSelection(apiClient *openapiclient.APIClient) (*HandleRouteSelectionResponse, error) {
	slog.Info("Handling route selection")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	locations, err := FetchStations(ctx, apiClient, "en")
	if err != nil {
		return nil, err
	}

	departingStation, err := selectDepartingStation(locations)
	if err != nil {
		return nil, err
	}

	arrivingStation, err := selectArrivingStation(locations)
	if err != nil {
		return nil, err
	}

	departureDate, err := selectDepartureDate()
	if err != nil {
		return nil, err
	}

	routes, err := FetchRoutes(apiClient, departingStation.StationID, arrivingStation.StationID, departureDate)
	if err != nil {
		return nil, err
	}

	selectedRoute, err := selectRoute(routes)
	if err != nil {
		return nil, err
	}

	response := &HandleRouteSelectionResponse{
		DepartingStation: departingStation,
		ArrivingStation:  arrivingStation,
		SelectedRoute:    selectedRoute,
	}

	return response, nil
}
