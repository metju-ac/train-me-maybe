package purchase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/internal/notification"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

// ```
// https://brn-ybus-pubapi.sa.cz/restapi/payments/credit/charge
// ```

var ErrFailedToPayTicket = errors.New("failed to pay ticket")

type Ticket struct {
	Type string `json:"type"`
	ID   int64  `json:"id"`
}

type FormField struct {
	FieldType  string `json:"fieldType"`
	FieldValue string `json:"fieldValue"`
}

// \"formFields\":[{\"fieldType\":\"EMAIL\",\"fieldValue\":\"robert.gemrot@centrum.cz\"}]}

type PayTicketRequest struct {
	Tickets    []Ticket    `json:"tickets"`
	FormFields []FormField `json:"formFields,omitempty"`
}

type PayTicketResponse struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func generateTxToken() string {
	const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	base := int64(len(base58Alphabet))
	timestamp := time.Now().UnixMilli()
	var encoded string

	// Generate the base58-encoded string
	for timestamp > 0 {
		remainder := timestamp % base
		timestamp /= base
		encoded = string(base58Alphabet[remainder]) + encoded
	}

	// Pad the result to a minimum length of 8
	for len(encoded) < 8 {
		encoded = "1" + encoded
	}

	return encoded
}

func payTicket(
	ctx context.Context,
	config *config.Config,
	input *models.UserInput,
	authToken string,
	ticket *openapiclient.Ticket,
	user *openapiclient.User,
) (*PayTicketResponse, error) {
	body := PayTicketRequest{
		Tickets: []Ticket{
			{
				Type: "RJ_SEAT",
				ID:   ticket.Id,
			},
		},
		FormFields: []FormField{
			{
				FieldType:  "EMAIL",
				FieldValue: *user.Email,
			},
		},
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		slog.Error("Failed to marshal ticket payment request", "error", err)
		return nil, fmt.Errorf("failed to marshal ticket payment request: %w", err)
	}

	buffer := bytes.NewBuffer(bodyJSON)

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, config.General.APIBaseURL+"/payments/credit/charge", buffer)
	if err != nil {
		slog.Error("Failed to create ticket payment request", "error", err)
		return nil, fmt.Errorf("failed to create ticket payment request: %w", err)
	}

	request.Header.Set("Host", "brn-ybus-pubapi.sa.cz")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:133.0) Gecko/20100101 Firefox/133.0")
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Accept-Language", "en-US,en;q=0.5")
	request.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	request.Header.Set("X-Lang", "cs")
	request.Header.Set("X-Currency", "CZK")
	request.Header.Set("Content-Type", "application/1.2.0+json")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("X-Application-Origin", "WEB")
	request.Header.Set("Authorization", "Bearer "+authToken)
	request.Header.Set("X-Txtoken", generateTxToken())
	request.Header.Set("Content-Length", "48")
	request.Header.Set("Origin", "https://regiojet.cz")
	request.Header.Set("DNT", "1")
	request.Header.Set("Sec-GPC", "1")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Referer", "https://regiojet.cz/")
	request.Header.Set("Sec-Fetch-Dest", "empty")
	request.Header.Set("Sec-Fetch-Mode", "cors")
	request.Header.Set("Sec-Fetch-Site", "cross-site")

	// protocol version is HTTP/2
	request.Proto = "HTTP/2"

	slog.Info("Sending ticket payment request", "url", request.URL.String())

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		slog.Error("Failed to send ticket payment request", "error", err)
		return nil, fmt.Errorf("failed to send ticket payment request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("Failed to pay ticket", "status", resp.Status)
		notification.EmailNotificationTicketNotPaid(&config.SMTP, input, ticket.Price, string(ticket.Currency))
		return nil, fmt.Errorf("%w", ErrFailedToPayTicket)
	}

	var response PayTicketResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		slog.Error("Failed to decode ticket payment response", "error", err)
		return nil, fmt.Errorf("failed to decode ticket payment response: %w", err)
	}

	return &response, nil
}
