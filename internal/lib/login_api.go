package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

var ErrFailedToLogin = errors.New("failed to login")

const contextTimeout = 30 * time.Second

type LoginRequest struct {
	AccountCode string `json:"accountCode"`
	Password    string `json:"password"`
}

func LoginWithCreditTicket(baseURL string, username string, password string) (string, error) {
	body := LoginRequest{
		AccountCode: username,
		Password:    password,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		slog.Error("Failed to marshal login request", "error", err)
		return "", fmt.Errorf("failed to marshal login request: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	buffer := bytes.NewBuffer(bodyJSON)

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/users/login/registeredAccount", buffer)
	if err != nil {
		slog.Error("Failed to create login request", "error", err)
		return "", fmt.Errorf("failed to create login request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")

	slog.Info("Sending login request", "url", request.URL.String())
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		slog.Error("Failed to send login request", "error", err)
		return "", fmt.Errorf("failed to send login request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("Failed to login", "status", resp.Status)
		return "", fmt.Errorf("%w: status %s", ErrFailedToLogin, resp.Status)
	}

	var response struct {
		Token string `json:"token"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		slog.Error("Failed to decode login response", "error", err)
		return "", fmt.Errorf("failed to decode login response: %w", err)
	}

	return response.Token, nil
}
