package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) GetStations(c *gin.Context) {
	lang := c.DefaultQuery("lang", "en")

	// Cache for 1 hour
	c.Header("Cache-Control", "max-age=3600, public") // Cache for 1 hour
	c.Header("Expires", time.Now().Add(1*time.Hour).Format(http.TimeFormat))

	stations, ok := h.LangStations[lang]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Language not supported"})
		return
	}

	c.JSON(http.StatusOK, stations)
}
