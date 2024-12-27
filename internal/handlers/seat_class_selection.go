package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/cli"
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

func HandleSeatClassSelection(apiClient *openapiclient.APIClient) ([]string, error) {
	slog.Info("Handling seat class selection")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	seatClasses, err := fetchSeatClasses(ctx, apiClient)
	if err != nil {
		return nil, err
	}

	selectedSeatClasses, err := selectSeatClasses(seatClasses)
	if err != nil {
		return nil, err
	}

	return selectedSeatClasses, nil
}
