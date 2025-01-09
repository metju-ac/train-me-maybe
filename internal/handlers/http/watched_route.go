package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/metju-ac/train-me-maybe/internal/dbmodels"
	"github.com/metju-ac/train-me-maybe/internal/handlers"
	"github.com/metju-ac/train-me-maybe/internal/lib"
	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/internal/notification"
	"github.com/metju-ac/train-me-maybe/internal/purchase"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type CreateWatchedRouteRequest struct {
	FromStationID       int64    `json:"fromStationId" binding:"required"`
	ToStationID         int64    `json:"toStationId" binding:"required"`
	RouteID             string   `json:"routeId" binding:"required"`
	SelectedSeatClasses []string `json:"selectedSeatClasses" binding:"required"`
	AutoPurchase        bool     `json:"autoPurchase"`
	TariffClass         string   `json:"tariffClass"`
	CreditUserNumber    string   `json:"creditUser"`
	CreditUserPassword  string   `json:"creditPassword"`
	CutOffTime          *int     `json:"cutOffTime"`
	MinimalCredit       *int     `json:"minimalCredit"`
}

func (h *Handler) cancelWatchedRoute(id uint) error {
	h.mu.Lock()
	if cancel, exists := h.GoroutineContexts[id]; exists {
		cancel()
		delete(h.GoroutineContexts, id)
	}
	h.mu.Unlock()

	err := h.WatchedRouteRepo.Delete(uint(id))
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
	response, err := purchase.AutoPurchaseTicket(h.ApiClient, &h.Config, userInput, checkFreeSeatsResponse)
	if err != nil {
		slog.Error("Error purchasing ticket", "error", err)
		_ = h.cancelWatchedRoute(watchedRoute.ID)
		return
	}

	slog.Info("Ticket purchased successfully", "watchedRouteID", watchedRoute.RouteID, "response", response)
	notification.EmailNotificationTicketBought(&h.Config.Smtp, userInput)

	if watchedRoute.MinimalCredit != nil && *watchedRoute.MinimalCredit >= int(response.Amount) {
		slog.Info("Low credit", "watchedRouteID", watchedRoute.RouteID)
		notification.EmailNotificationLowCredit(&h.Config.Smtp, userInput, response.Amount, response.Currency)
	}

	_ = h.cancelWatchedRoute(watchedRoute.ID)
}

func (h *Handler) CreateWatchedRoute(c *gin.Context) {
	var req CreateWatchedRouteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Invalid request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if req.AutoPurchase && (req.CreditUserNumber == "" || req.CreditUserPassword == "" || req.TariffClass == "") {
		slog.Error("Missing required fields for auto purchase")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields for auto purchase"})
		return
	}

	email, exists := c.Get("email")
	if !exists {
		slog.Warn("Unauthorized access attempt")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if req.AutoPurchase {
		_, err := lib.LoginWithCreditTicket(h.Config.General.ApiBaseUrl, req.CreditUserNumber, req.CreditUserPassword)
		if err != nil {
			slog.Error("Failed to login with credit ticket", "error", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to login with credit ticket"})
			return
		}
	}

	watchedRoute := &dbmodels.WatchedRoute{
		UserEmail:           email.(string),
		FromStationID:       req.FromStationID,
		ToStationID:         req.ToStationID,
		RouteID:             req.RouteID,
		SelectedSeatClasses: req.SelectedSeatClasses,
		AutoPurchase:        req.AutoPurchase,
		TariffClass:         req.TariffClass,
		CreditUserNumber:    req.CreditUserNumber,
		CreditUserPassword:  req.CreditUserPassword,
		CutOffTime:          req.CutOffTime,
		MinimalCredit:       req.MinimalCredit,
	}

	if err := h.WatchedRouteRepo.Create(watchedRoute); err != nil {
		slog.Error("Error creating watched route", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating watched route"})
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	h.mu.Lock()
	h.GoroutineContexts[watchedRoute.ID] = cancel
	h.mu.Unlock()

	go func() {
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
	}()

	slog.Info("Watched route created successfully", "id", watchedRoute.ID)
	c.JSON(http.StatusOK, watchedRoute)
}

func (h *Handler) GetWatchedRoutes(c *gin.Context) {
	mail, exists := c.Get("email")
	if !exists {
		slog.Warn("Unauthorized access attempt")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	routes, err := h.WatchedRouteRepo.GetAllForEmail(mail.(string))
	if err != nil {
		slog.Error("Error retrieving watched routes", "email", mail, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving watched routes"})
		return
	}

	slog.Info("Watched routes retrieved successfully", "email", mail)
	c.JSON(http.StatusOK, routes)
}

func (h *Handler) DeleteWatchedRoute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		slog.Error("Invalid route ID", "id", idParam, "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID"})
		return
	}

	if err := h.cancelWatchedRoute(uint(id)); err != nil {
		slog.Error("Error deleting watched route", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting watched route"})
		return
	}

	slog.Info("Watched route deleted successfully", "id", id)
	c.JSON(http.StatusOK, gin.H{"message": "Watched route deleted successfully"})
}
