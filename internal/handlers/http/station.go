package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) GetStations(c *gin.Context) {
	// Cache for 1 hour
	c.Header("Cache-Control", "max-age=3600, public") // Cache for 1 hour
	c.Header("Expires", time.Now().Add(1*time.Hour).Format(http.TimeFormat))

	c.JSON(http.StatusOK, h.Stations)
}
