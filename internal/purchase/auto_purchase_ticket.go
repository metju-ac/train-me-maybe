package purchase

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/handlers"
	"github.com/metju-ac/train-me-maybe/internal/lib"
	"github.com/metju-ac/train-me-maybe/internal/models"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

const contextTimeout = 30 * time.Second

func getUserDetails(ctx context.Context, apiClient *openapiclient.APIClient, token string) (*openapiclient.User, error) {
	slog.Info("Getting details about the user")

	user, httpRes, err := apiClient.UsersAPI.Authenticate(ctx).Authorization(token).Execute()
	if err != nil {
		slog.Error("Failed to get details about the user", "error", err)
		return nil, fmt.Errorf("failed to get details about the user: %w", err)
	}
	defer httpRes.Body.Close()
	slog.Info("Successfully got details about the user", "statusCode", httpRes.StatusCode)

	return user, nil
}

func AutoPurchaseTicket(
	apiClient *openapiclient.APIClient,
	config *config.Config,
	input *models.UserInput,
	seats *handlers.CheckFreeSeatsResponse,
) (*PayTicketResponse, *openapiclient.Ticket, error) {
	if !config.Auth.CreditEnabled {
		return nil, nil, nil
	}

	slog.Info("Beginning with auto purchasing ticket")

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	token, err := lib.LoginWithCreditTicket(config.General.APIBaseURL, input.CreditUserNumber, input.CreditUserPassword)
	if err != nil {
		slog.Error("Failed to login with credit ticket", "error", err)
		return nil, nil, fmt.Errorf("failed to login with credit ticket: %w", err)
	}

	user, err := getUserDetails(ctx, apiClient, token)
	if err != nil {
		slog.Error("Failed to get user details", "error", err)
		return nil, nil, fmt.Errorf("failed to get user details: %w", err)
	}

	ticket, err := createTicket(ctx, config, input, seats, user, token)
	if err != nil {
		slog.Error("Failed to create ticket", "error", err)
		return nil, nil, fmt.Errorf("failed to create ticket: %w", err)
	}

	payment, err := payTicket(ctx, config, input, token, ticket, user)
	if err != nil {
		slog.Error("Failed to pay ticket", "error", err)
		return nil, nil, fmt.Errorf("failed to pay ticket: %w", err)
	}

	slog.Info("Successfully purchased ticket", "amount", payment.Amount, "currency", payment.Currency)

	return payment, ticket, nil
}
