package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/models"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func FetchRouteDetail(apiClient *openapiclient.APIClient, input *models.UserInput) (*openapiclient.Route, error) {
	slog.Info("Fetching route detail")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	routeDetail, httpRes, err := apiClient.RoutesAPI.GetSimpleRouteDetail(ctx, input.SelectedRoute.Id).FromStationId(input.DepartingStation.StationID).ToStationId(input.ArrivingStation.StationID).Execute()
	if err != nil {
		slog.Error("Failed to fetch route detail", "error", err)
		return nil, fmt.Errorf("failed to fetch route detail: %w", err)
	}
	slog.Info("Successfully fetched route detail", "statusCode", httpRes.StatusCode)
	return routeDetail, nil
}
