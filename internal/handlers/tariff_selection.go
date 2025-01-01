package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/cli"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func fetchTariffs(ctx context.Context, apiClient *openapiclient.APIClient) ([]openapiclient.Tariff, error) {
	slog.Info("Fetching tariffs")
	tariffs, httpRes, err := apiClient.ConstsAPI.GetTariffs(ctx).Execute()
	if err != nil {
		slog.Error("Failed to fetch tariffs", "error", err)
		return nil, fmt.Errorf("failed to fetch tariffs: %w", err)
	}

	slog.Info("Successfully fetched tariffs", "statusCode", httpRes.StatusCode)
	return tariffs, nil
}

func selectTariff(tariff []openapiclient.Tariff) (*openapiclient.Tariff, error) {
	slog.Info("Selecting tariffs")

	selectedTariff, err := cli.SelectTariff(tariff)
	if err != nil {
		slog.Error("Failed to select tariffs", "error", err)
		return nil, fmt.Errorf("failed to select tariffs: %w", err)
	}

	slog.Info("Successfully selected tariffs", "selectedTariff", selectedTariff)
	return selectedTariff, nil
}

func DefaultTariff() *openapiclient.Tariff {
	regular := "REGULAR"
	adult := "Adult 18-64 years"
	return &openapiclient.Tariff{
		Key:   &regular,
		Value: &adult,
	}
}

func HandleTariffSelection(apiClient *openapiclient.APIClient) (*openapiclient.Tariff, error) {
	slog.Info("Handling tariff selection")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tariffs, err := fetchTariffs(ctx, apiClient)
	if err != nil {
		return nil, err
	}

	selectedTariff, err := selectTariff(tariffs)
	if err != nil {
		return nil, err
	}

	return selectedTariff, nil
}
