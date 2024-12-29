package purchase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"

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
			RouteId:     input.SelectedRoute.Id,
			SeatClass:   seats.SeatClass,
			PriceSource: input.RouteDetail.PriceClasses[0].PriceSource,
			ActionPrice: nil,
			Sections: []openapiclient.CreateTicketSectionRequest{
				{
					Section: *input.Section,
					SelectedSeats: []openapiclient.SelectedSeat100{openapiclient.SelectedSeat100{
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
				Tariff: *input.Tariff.Key,
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

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+authToken)

	slog.Info("Sending ticket creation request", "url", request.URL.String())

	// dump body to file ticket-creation.json
	os.WriteFile("ticket-creation.json", bodyJson, 0644)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		slog.Error("Failed to send ticket creation request", "error", err)
		return nil, err
	}

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
