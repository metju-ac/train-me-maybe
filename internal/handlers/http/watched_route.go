package http

import (
	"github.com/gin-gonic/gin"
	"github.com/metju-ac/train-me-maybe/internal/dbmodels"
	"log/slog"
	"net/http"
	"strconv"
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
	}

	if err := h.WatchedRouteRepo.Create(watchedRoute); err != nil {
		slog.Error("Error creating watched route", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating watched route"})
		return
	}

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

	err = h.WatchedRouteRepo.Delete(uint(id))
	if err != nil {
		slog.Error("Error deleting watched route", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting watched route"})
		return
	}

	slog.Info("Watched route deleted successfully", "id", id)
	c.JSON(http.StatusOK, gin.H{"message": "Watched route deleted successfully"})
}
