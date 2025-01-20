package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metju-ac/train-me-maybe/internal/dbmodels"
	"github.com/metju-ac/train-me-maybe/internal/lib"
)

type CreateWatchedRouteRequest struct {
	FromStationID       int64    `binding:"required"    json:"fromStationId"`
	ToStationID         int64    `binding:"required"    json:"toStationId"`
	RouteID             string   `binding:"required"    json:"routeId"`
	SelectedSeatClasses []string `binding:"required"    json:"selectedSeatClasses"`
	AutoPurchase        bool     `json:"autoPurchase"`
	TariffClass         string   `json:"tariffClass"`
	CreditUserNumber    string   `json:"creditUser"`
	CreditUserPassword  string   `json:"creditPassword"`
	CutOffTime          *int     `json:"cutOffTime"`
	MinimalCredit       *int     `json:"minimalCredit"`
	DepartureDateTime   string   `binding:"required"    json:"departureDateTime"`
}

func (h *Handler) CreateWatchedRoute(c *gin.Context) {
	var req CreateWatchedRouteRequest
	fmt.Println(c)
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

	emailStr, ok := email.(string)
	if !ok {
		slog.Error("Email type assertion failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if req.AutoPurchase {
		_, err := lib.LoginWithCreditTicket(h.Config.General.APIBaseURL, req.CreditUserNumber, req.CreditUserPassword)
		if err != nil {
			slog.Error("Failed to login with credit ticket", "error", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to login with credit ticket"})
			return
		}
	}

	departureDateTime, err := time.Parse(time.RFC3339, req.DepartureDateTime)
	if err != nil {
		slog.Error("Invalid departureDateTime format", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid departureDateTime format"})
		return
	}

	watchedRoute := &dbmodels.WatchedRoute{
		UserEmail:           emailStr,
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
		DepartureDateTime:   departureDateTime,
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

	go h.startWatchingRoute(ctx, watchedRoute)

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
