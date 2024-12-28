package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/cli"
	"github.com/metju-ac/train-me-maybe/internal/lib"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func fetchSeatClasses(ctx context.Context, apiClient *openapiclient.APIClient) ([]openapiclient.SeatClass, error) {
	slog.Info("Fetching seat classes")
	seatClasses, httpRes, err := apiClient.ConstsAPI.GetSeatClasses(ctx).Execute()
	if err != nil {
		slog.Error("Failed to fetch seat classes", "error", err)
		return nil, fmt.Errorf("failed to fetch seat classes: %w", err)
	}

	slog.Info("Successfully fetched seat classes", "statusCode", httpRes.StatusCode)
	return seatClasses, nil
}

func fetchSeatClassesInRoute(ctx context.Context, apiClient *openapiclient.APIClient, route *HandleRouteSelectionResponse) ([]openapiclient.SeatClass, error) {
	slog.Info("Fetching route detail")
	routeDetail, httpRes, err := apiClient.RoutesAPI.GetSimpleRouteDetail(ctx, route.SelectedRoute.Id).FromStationId(route.DepartingStation.StationID).ToStationId(route.ArrivingStation.StationID).Execute()

	if err != nil {
		slog.Error("Failed to fetch route detail", "error", err)
		return nil, fmt.Errorf("failed to fetch route detail: %w", err)
	}

	slog.Info("Successfully fetched route detail", "statusCode", httpRes.StatusCode)

	// create a "set" to find out if a seat class is present
	presentSeatClasses := lib.MapFunc(routeDetail.PriceClasses, func(r openapiclient.PriceClass) string { return r.SeatClassKey })
	presentSeatClassesDict := lib.ToSet(presentSeatClasses)

	seatClasses, err := fetchSeatClasses(ctx, apiClient)
	if err != nil {
		return nil, err
	}

	result := lib.FilterFunc(seatClasses, func(r openapiclient.SeatClass) bool { return presentSeatClassesDict[r.Key] })

	return result, nil
}

func selectSeatClasses(seatClasses []openapiclient.SeatClass) ([]string, error) {
	slog.Info("Selecting seat classes")

	selectedSeatClasses, err := cli.SelectSeatClasses(seatClasses)
	if err != nil {
		slog.Error("Failed to select seat classes", "error", err)
		return nil, fmt.Errorf("failed to select seat classes: %w", err)
	}

	slog.Info("Successfully selected seat classes", "selectedSeatClasses", selectedSeatClasses)
	return selectedSeatClasses, nil
}

func HandleSeatClassSelection(apiClient *openapiclient.APIClient, route *HandleRouteSelectionResponse) ([]string, error) {
	slog.Info("Handling seat class selection")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	seatClasses, err := fetchSeatClassesInRoute(ctx, apiClient, route)
	if err != nil {
		return nil, err
	}

	selectedSeatClasses, err := selectSeatClasses(seatClasses)
	if err != nil {
		return nil, err
	}

	return selectedSeatClasses, nil
}
