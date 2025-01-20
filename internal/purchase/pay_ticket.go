package purchase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/internal/notification"

	"github.com/metju-ac/train-me-maybe/internal/config"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

// ```
// https://brn-ybus-pubapi.sa.cz/restapi/payments/credit/charge
// ```

type Ticket struct {
	Type string `json:"type"`
	Id   int64  `json:"id"`
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

func payTicket(ctx context.Context, config *config.Config, input *models.UserInput, authToken string, ticket *openapiclient.Ticket, user *openapiclient.User) (*PayTicketResponse, error) {
	body := PayTicketRequest{
		Tickets: []Ticket{
			{
				Type: "RJ_SEAT",
				Id:   ticket.Id,
			},
		},
		FormFields: []FormField{
			{
				FieldType:  "EMAIL",
				FieldValue: *user.Email,
			},
		},
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		slog.Error("Failed to marshal ticket payment request", "error", err)
		return nil, err
	}

	buffer := bytes.NewBuffer(bodyJson)

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, config.General.ApiBaseUrl+"/payments/credit/charge", buffer)
	if err != nil {
		slog.Error("Failed to create ticket payment request", "error", err)
		return nil, err
	}

	// set these headers:
	// {
	// 	"name": "Host",
	// 	"value": "brn-ybus-pubapi.sa.cz"
	//   },
	//   {
	// 	"name": "User-Agent",
	// 	"value": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:133.0) Gecko/20100101 Firefox/133.0"
	//   },
	//   {
	// 	"name": "Accept",
	// 	"value": "application/json, text/plain, */*"
	//   },
	//   {
	// 	"name": "Accept-Language",
	// 	"value": "en-US,en;q=0.5"
	//   },
	//   {
	// 	"name": "Accept-Encoding",
	// 	"value": "gzip, deflate, br, zstd"
	//   },
	//   {
	// 	"name": "X-Lang",
	// 	"value": "cs"
	//   },
	//   {
	// 	"name": "X-Currency",
	// 	"value": "CZK"
	//   },
	//   {
	// 	"name": "Content-Type",
	// 	"value": "application/1.2.0+json"
	//   },
	//   {
	// 	"name": "Cache-Control",
	// 	"value": "no-cache"
	//   },
	//   {
	// 	"name": "X-Application-Origin",
	// 	"value": "WEB"
	//   },
	//   {
	// 	"name": "Authorization",
	// 	"value": "Bearer 8a64a1c1-6611-45ea-9ff6-15919d8ebd5f"
	//   },
	//   {
	// 	"name": "X-TxToken",
	// 	"value": "182qJ8bn"
	//   },
	//   {
	// 	"name": "Content-Length",
	// 	"value": "48"
	//   },
	//   {
	// 	"name": "Origin",
	// 	"value": "https://regiojet.cz"
	//   },
	//   {
	// 	"name": "DNT",
	// 	"value": "1"
	//   },
	//   {
	// 	"name": "Sec-GPC",
	// 	"value": "1"
	//   },
	//   {
	// 	"name": "Connection",
	// 	"value": "keep-alive"
	//   },
	//   {
	// 	"name": "Referer",
	// 	"value": "https://regiojet.cz/"
	//   },
	//   {
	// 	"name": "Sec-Fetch-Dest",
	// 	"value": "empty"
	//   },
	//   {
	// 	"name": "Sec-Fetch-Mode",
	// 	"value": "cors"
	//   },
	//   {
	// 	"name": "Sec-Fetch-Site",
	// 	"value": "cross-site"
	//   }

	request.Header.Set("Host", "brn-ybus-pubapi.sa.cz")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:133.0) Gecko/20100101 Firefox/133.0")
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Accept-Language", "en-US,en;q=0.5")
	request.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	request.Header.Set("X-Lang", "cs")
	request.Header.Set("X-Currency", "CZK") // TODO: get currency from the user, and fallback to CZK
	request.Header.Set("Content-Type", "application/1.2.0+json")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("X-Application-Origin", "WEB")
	request.Header.Set("Authorization", "Bearer "+authToken)
	request.Header.Set("X-TxToken", generateTxToken())
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

	// dump body to file ticket-payment.json
	os.WriteFile("ticket-payment.json", bodyJson, 0o644)

	dump, err := httputil.DumpRequestOut(request, true)
	if err != nil {
		return nil, err
	}
	os.WriteFile("ticket-payment-request.txt", dump, 0o644)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		slog.Error("Failed to send ticket payment request", "error", err)
		return nil, err
	}

	dump, err = httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}
	os.WriteFile("ticket-payment-response.txt", dump, 0o644)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("Failed to pay ticket", "status", resp.Status)
		notification.EmailNotificationTicketNotPaid(&config.Smtp, input, ticket.Price, string(ticket.Currency))
		return nil, errors.New("Failed to pay ticket")
	}

	var response PayTicketResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		slog.Error("Failed to decode ticket payment response", "error", err)
		return nil, err
	}

	return &response, nil
}
