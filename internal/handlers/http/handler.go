package http

import (
	"context"
	"fmt"
	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/dbmodels"
	"github.com/metju-ac/train-me-maybe/internal/handlers"
	"github.com/metju-ac/train-me-maybe/internal/lib"
	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/internal/notification"
	"github.com/metju-ac/train-me-maybe/internal/purchase"
	"github.com/metju-ac/train-me-maybe/internal/repositories"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
	"log/slog"
	"sync"
	"time"
)

type Handler struct {
	Stations               map[int64]models.StationModel
	UserRepo               repositories.UserRepository
	WatchedRouteRepo       repositories.WatchedRouteRepository
	SuccessfulPurchaseRepo repositories.SuccessfulPurchaseRepository
	ApiClient              *openapiclient.APIClient
	GoroutineContexts      map[uint]context.CancelFunc
	mu                     sync.Mutex
	Config                 config.Config
}

func (h *Handler) cancelWatchedRoute(id uint) error {
	h.mu.Lock()
	if cancel, exists := h.GoroutineContexts[id]; exists {
		cancel()
		delete(h.GoroutineContexts, id)
	}
	h.mu.Unlock()

	err := h.WatchedRouteRepo.Delete(id)
	if err != nil {
		slog.Error("Error deleting watched route", "id", id, "error", err)
		return err
	}

	return nil
}

func (h *Handler) getUserInput(watchedRoute *dbmodels.WatchedRoute) (*models.UserInput, error) {
	departingStation, exist := h.Stations[watchedRoute.FromStationID]
	if !exist {
		slog.Error("Departing station not found", "id", watchedRoute.FromStationID)
		return nil, fmt.Errorf("Departing station not found")
	}

	arrivingStation, exist := h.Stations[watchedRoute.ToStationID]
	if !exist {
		slog.Error("Arriving station not found", "id", watchedRoute.ToStationID)
		return nil, fmt.Errorf("Arriving station not found")
	}

	userInput := &models.UserInput{
		DepartingStation:   &departingStation,
		ArrivingStation:    &arrivingStation,
		SelectedRouteIds:   watchedRoute.RouteID,
		SeatClasses:        watchedRoute.SelectedSeatClasses,
		TariffKey:          watchedRoute.TariffClass,
		Section:            nil,
		RouteDetail:        nil,
		CreditUserNumber:   watchedRoute.CreditUserNumber,
		CreditUserPassword: watchedRoute.CreditUserPassword,
		UserEmail:          watchedRoute.UserEmail,
	}

	routeDetail, err := handlers.FetchRouteDetail(h.ApiClient, userInput)
	if err != nil {
		slog.Error("Error fetching route detail", "error", err)
		return nil, fmt.Errorf("Error fetching route detail: %w", err)
	}
	userInput.RouteDetail = routeDetail

	section, err := lib.GetSection(userInput)
	if err != nil {
		slog.Error("Error getting section", "error", err)
		return nil, fmt.Errorf("Error getting section: %w", err)
	}
	userInput.Section = section

	return userInput, nil
}

func (h *Handler) watchRoute(userInput *models.UserInput, watchedRoute *dbmodels.WatchedRoute) {
	slog.Error("Watching route", "watchedRouteID", watchedRoute.ID)
	if time.Now().After(userInput.RouteDetail.DepartureTime) {
		slog.Info("Departure time passed", "watchedRouteID", watchedRoute.RouteID)
		_ = h.cancelWatchedRoute(watchedRoute.ID)
		return
	}

	checkFreeSeatsResponse, err := handlers.CheckFreeSeats(h.ApiClient, userInput)
	if err != nil {
		slog.Error("Error checking free seats", "error", err)
		return
	}

	if checkFreeSeatsResponse == nil || !checkFreeSeatsResponse.HasFreeSeats {
		slog.Info("No free seats found", "routeID", watchedRoute.RouteID)
		return
	}
	slog.Info("Free seats found", "routeID", watchedRoute.RouteID)

	cutOffTimeDuration := time.Duration(0)
	if watchedRoute.CutOffTime != nil {
		cutOffTimeDuration = time.Duration(*watchedRoute.CutOffTime) * time.Minute
	}

	if !watchedRoute.AutoPurchase || !time.Now().Add(cutOffTimeDuration).Before(userInput.RouteDetail.DepartureTime) {
		slog.Info("Auto purchase disabled or cut off time passed", "watchedRouteID", watchedRoute.RouteID)
		notification.EmailNotificationFreeSeats(&h.Config.Smtp, userInput)
		_ = h.cancelWatchedRoute(watchedRoute.ID)
		return
	}

	slog.Info("Purchasing ticket", "watchedRouteID", watchedRoute.RouteID)
	response, ticket, err := purchase.AutoPurchaseTicket(h.ApiClient, &h.Config, userInput, checkFreeSeatsResponse)
	if err != nil {
		slog.Error("Error purchasing ticket", "error", err)
		_ = h.cancelWatchedRoute(watchedRoute.ID)
		return
	}

	slog.Info("Ticket purchased successfully", "watchedRouteID", watchedRoute.RouteID, "response", response)

	if watchedRoute.MinimalCredit != nil && *watchedRoute.MinimalCredit >= int(response.Amount) {
		slog.Info("Low credit", "watchedRouteID", watchedRoute.RouteID)
		notification.EmailNotificationLowCredit(&h.Config.Smtp, userInput, response.Amount, response.Currency)
	}

	successfulPurchase := &dbmodels.SuccessfulPurchase{
		UserEmail:           watchedRoute.UserEmail,
		FromStationID:       watchedRoute.FromStationID,
		ToStationID:         watchedRoute.ToStationID,
		RouteID:             watchedRoute.RouteID,
		TariffClass:         watchedRoute.TariffClass,
		SelectedSeatClasses: watchedRoute.SelectedSeatClasses,
		PurchaseTime:        time.Now(),
		Price:               ticket.Price,
		Currency:            string(ticket.Currency),
		SeatClass:           ticket.SeatClassKey,
	}

	beersOwed, err := h.SuccessfulPurchaseRepo.CreateWithBeerUpdate(successfulPurchase, h.Config.General.TicketBeerPrice)
	if err != nil {
		slog.Error("Error creating successful purchase", "error", err)
	}

	notification.EmailNotificationTicketBoughtWithBeers(&h.Config.Smtp, userInput, ticket, beersOwed)
	_ = h.cancelWatchedRoute(watchedRoute.ID)
}

func (h *Handler) startWatchingRoute(ctx context.Context, watchedRoute *dbmodels.WatchedRoute) {
	ticker := time.NewTicker(time.Duration(h.Config.General.PollInterval) * time.Second)
	defer ticker.Stop()

	userInput, err := h.getUserInput(watchedRoute)
	if err != nil {
		slog.Error("Error getting user input", "error", err)
		_ = h.cancelWatchedRoute(watchedRoute.ID)
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			h.watchRoute(userInput, watchedRoute)
		}
	}
}

func (h *Handler) RestartWatchedRoutes() error {
	routes, err := h.WatchedRouteRepo.GetAll()
	if err != nil {
		slog.Error("Error retrieving watched routes", "error", err)
		return err
	}

	for _, watchedRoute := range routes {
		ctx, cancel := context.WithCancel(context.Background())
		h.mu.Lock()
		h.GoroutineContexts[watchedRoute.ID] = cancel
		h.mu.Unlock()

		go h.startWatchingRoute(ctx, &watchedRoute)
	}

	slog.Info("Restarted watching all routes")
	return nil
}
