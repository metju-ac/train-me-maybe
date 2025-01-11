package http

import (
	"github.com/gin-gonic/gin"
	"github.com/metju-ac/train-me-maybe/internal/models"
	"net/http"
	"time"
)

func (h *Handler) GetStations(c *gin.Context) {
	// Cache for 1 hour
	c.Header("Cache-Control", "max-age=3600, public") // Cache for 1 hour
	c.Header("Expires", time.Now().Add(1*time.Hour).Format(http.TimeFormat))

	stations := make([]models.StationModel, 0, len(h.Stations))
	for _, station := range h.Stations {
		stations = append(stations, station)
	}

	c.JSON(http.StatusOK, stations)
}
