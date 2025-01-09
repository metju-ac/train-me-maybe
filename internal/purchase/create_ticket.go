package purchase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/handlers"
	"github.com/metju-ac/train-me-maybe/internal/models"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

type CreateRegisteredTicketRequest struct {
	TicketRequests []openapiclient.CreateTicketRequest `json:"ticketRequests"`
}

func createTicket(ctx context.Context, config *config.Config, input *models.UserInput, seats *handlers.CheckFreeSeatsResponse, user *openapiclient.User, authToken string) (*openapiclient.Ticket, error) {
	payload := openapiclient.CreateTicketRequest{
		Route: openapiclient.CreateTicketRouteRequest{
			RouteId:     input.SelectedRouteIds,
			SeatClass:   seats.SeatClass,
			PriceSource: input.RouteDetail.PriceClasses[0].PriceSource,
			ActionPrice: nil,
			Sections: []openapiclient.CreateTicketSectionRequest{
				{
					Section: *input.Section,
					SelectedSeats: []openapiclient.SelectedSeat100{
						{
							SectionId:     seats.SelectedSeat.SectionId,
							VehicleNumber: seats.SelectedSeat.VehicleNumber,
							SeatIndex:     seats.SelectedSeat.SeatIndex,
						}},
				},
			},
		},
		SelectedAddons:        []openapiclient.SelectedAddon{},
		CodeDiscount:          nil,
		PercentualDiscountIds: []int64{},
		Passengers: []openapiclient.Passenger{
			{
				Tariff: input.TariffKey,
				Phone:  user.PhoneNumber,
				Email:  user.Email,
			},
		},
	}

	body := CreateRegisteredTicketRequest{
		TicketRequests: []openapiclient.CreateTicketRequest{payload},
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		slog.Error("Failed to marshal ticket creation request", "error", err)
		return nil, err
	}

	buffer := bytes.NewBuffer(bodyJson)

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, config.General.ApiBaseUrl+"/tickets/create/registered", buffer)
	if err != nil {
		slog.Error("Failed to create ticket creation request", "error", err)
		return nil, err
	}

	// Headers to be sent with the request (copied from firefox dev tools)

	//   {
	// 	  "name": "Host",
	// 	  "value": "brn-ybus-pubapi.sa.cz"
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
	// 	"value": "Bearer 5e2395a7-3d88-430f-ae0a-4b25e9f529f0"
	//   },
	//   {
	// 	"name": "Content-Length",
	// 	"value": "458"
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
	//   },
	//   {
	// 	"name": "TE",
	// 	"value": "trailers"
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
	request.Header.Set("Content-Length", "458")
	request.Header.Set("Origin", "https://regiojet.cz")
	request.Header.Set("DNT", "1")
	request.Header.Set("Sec-GPC", "1")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Referer", "https://regiojet.cz/")
	request.Header.Set("Sec-Fetch-Dest", "empty")
	request.Header.Set("Sec-Fetch-Mode", "cors")
	request.Header.Set("Sec-Fetch-Site", "cross-site")
	request.Header.Set("TE", "trailers")

	// protocol version is HTTP/2
	request.Proto = "HTTP/2"

	slog.Info("Sending ticket creation request", "url", request.URL.String())

	// dump body to file ticket-creation.json
	// os.WriteFile("ticket-creation.json", bodyJson, 0644)

	// dump, err := httputil.DumpRequestOut(request, true)
	// if err != nil {
	// 	return nil, err
	// }
	// os.WriteFile("ticket-creation-request.txt", dump, 0644)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		slog.Error("Failed to send ticket creation request", "error", err)
		return nil, err
	}

	// dump, err = httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	return nil, err
	// }
	// os.WriteFile("ticket-creation-response.txt", dump, 0644)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("Failed to create ticket", "status", resp.Status)
		return nil, errors.New("Failed to create ticket")
	}

	var response struct {
		Tickets []openapiclient.Ticket `json:"tickets"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		slog.Error("Failed to decode ticket creation response", "error", err)
		return nil, err
	}

	if len(response.Tickets) != 1 {
		slog.Error("Unexpected number of tickets created", "count", len(response.Tickets))
		return nil, errors.New("Unexpected number of tickets created")
	}

	return &response.Tickets[0], nil
}
