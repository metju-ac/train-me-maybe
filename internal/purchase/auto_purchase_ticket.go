package purchase

import (
	"context"
	"log/slog"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/handlers"
	"github.com/metju-ac/train-me-maybe/internal/lib"
	"github.com/metju-ac/train-me-maybe/internal/models"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func getUserDetails(ctx context.Context, apiClient *openapiclient.APIClient, token string) (*openapiclient.User, error) {
	slog.Info("Getting details about the user")

	user, httpRes, err := apiClient.UsersAPI.Authenticate(ctx).Authorization(token).Execute()
	if err != nil {
		slog.Error("Failed to get details about the user", "error", err)
		return nil, err
	}
	slog.Info("Successfully got details about the user", "statusCode", httpRes.StatusCode)

	return user, nil
}

func AutoPurchaseTicket(apiClient *openapiclient.APIClient, config *config.Config, input *models.UserInput, seats *handlers.CheckFreeSeatsResponse) error {
	if !config.Auth.CreditEnabled {
		return nil
	}

	slog.Info("Beginning with auto purchasing ticket")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	token, err := lib.LoginWithCreditTicket(ctx, config.General.ApiBaseUrl, config.Auth.CreditUser, config.Auth.CreditPassword)
	if err != nil {
		slog.Error("Failed to login with credit ticket", "error", err)
		return err
	}

	user, err := getUserDetails(ctx, apiClient, token)
	if err != nil {
		slog.Error("Failed to get user details", "error", err)
		return err
	}

	ticket, err := createTicket(ctx, config, input, seats, user, token)
	if err != nil {
		slog.Error("Failed to create ticket", "error", err)
		return err
	}

	slog.Info("Successfully created ticket", "ticketId", ticket.Id)

	return nil
}
