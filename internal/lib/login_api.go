package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type LoginRequest struct {
	AccountCode string `json:"accountCode"`
	Password    string `json:"password"`
}

func LoginWithCreditTicket(ctx context.Context, baseUrl string, username string, password string) (string, error) {
	body := LoginRequest{
		AccountCode: username,
		Password:    password,
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		slog.Error("Failed to marshal login request", "error", err)
		return "", err
	}

	buffer := bytes.NewBuffer(bodyJson)

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, baseUrl+"/users/login/registeredAccount", buffer)
	if err != nil {
		slog.Error("Failed to create login request", "error", err)
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")

	slog.Info("Sending login request", "url", request.URL.String())
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		slog.Error("Failed to send login request", "error", err)
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("Failed to login", "status", resp.Status)
		return "", errors.New("Failed to login")
	}

	var response struct {
		Token string `json:"token"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		slog.Error("Failed to decode login response", "error", err)
		return "", err
	}

	return response.Token, nil
}
