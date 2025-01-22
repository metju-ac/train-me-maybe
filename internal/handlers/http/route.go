package http

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metju-ac/train-me-maybe/internal/handlers"
	"github.com/metju-ac/train-me-maybe/internal/lib"
	"github.com/metju-ac/train-me-maybe/openapi"
)

const HoursInDay = 24

func (h *Handler) GetRoutes(c *gin.Context) {
	fromStationIDStr := c.Query("fromStationId")
	toStationIDStr := c.Query("toStationId")
	date := c.Query("date")

	fromStationID, err := strconv.ParseInt(fromStationIDStr, 10, 64)
	if err != nil {
		slog.Error("Invalid fromStationId", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fromStationId"})
		return
	}

	toStationID, err := strconv.ParseInt(toStationIDStr, 10, 64)
	if err != nil {
		slog.Error("Invalid toStationId", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid toStationId"})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		slog.Error("Invalid date format", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	if parsedDate.Before(time.Now().Truncate(HoursInDay * time.Hour)) {
		slog.Error("Date is in the past", "date", date)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date cannot be in the past"})
		return
	}

	routes, err := handlers.FetchRoutes(h.APIClient, fromStationID, toStationID, date)
	if err != nil {
		slog.Error("Failed to fetch routes", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch routes"})
		return
	}

	filteredRoutes := lib.FilterFunc(routes, func(route openapi.SimpleRoute) bool {
		return *route.TransfersCount == 0
	})

	slog.Info("Successfully fetched routes", "fromStationID", fromStationID, "toStationID", toStationID, "date", date)

	// Cache for 1 hour
	c.Header("Cache-Control", "max-age=3600, public")
	c.Header("Expires", time.Now().Add(1*time.Hour).Format(http.TimeFormat))

	c.JSON(http.StatusOK, filteredRoutes)
}
