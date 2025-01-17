package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) GetSeatClasses(c *gin.Context) {
	fromStationIDStr := c.Query("fromStationId")
	toStationIDStr := c.Query("toStationId")
	routeIdStr := c.Query("routeId")

	fromStationID, err := strconv.ParseInt(fromStationIDStr, 10, 64)
	if err != nil {
		slog.Error("Invalid fromStationId", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fromStationId"})
		return
	}

	toStationId, err := strconv.ParseInt(toStationIDStr, 10, 64)
	if err != nil {
		slog.Error("Invalid toStationId", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid toStationId"})
		return
	}

	slog.Info("Fetching route detail", "fromStationId", fromStationID, "toStationId", toStationId, "routeId", routeIdStr)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	routeDetail, httpRes, err := h.ApiClient.RoutesAPI.GetSimpleRouteDetail(ctx, routeIdStr).
		FromStationId(fromStationID).
		ToStationId(toStationId).
		Execute()
	if err != nil {
		slog.Error("Failed to fetch route detail", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching route detail"})
		return
	}
	slog.Info("Successfully fetched route detail", "statusCode", httpRes.StatusCode)

	if len(routeDetail.Sections) != 1 {
		slog.Error("Expected exactly one section in route detail", "sections", routeDetail.Sections)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Expected exactly one section in route detail"})
		return
	}

	vehicleStandardKey := routeDetail.Sections[0].VehicleStandardKey
	var seatClasses []string

	switch vehicleStandardKey {
	case "YELLOW":
		seatClasses = []string{"TRAIN_LOW_COST", "C0", "C1", "C2"}
	case "FUN_RELAX_SELF_SERVICE":
		seatClasses = []string{"NO"}
	case "FUN_RELAX_PLUS":
		seatClasses = []string{"BUS_STANDARD", "BUS_RELAX"}
	case "YELLOW_R8":
		lineGroupCode := *routeDetail.Sections[0].Line.LineGroupCode
		switch lineGroupCode {
		case "VLAK_R8":
			seatClasses = []string{"TRAIN_2ND_CLASS", "TRAIN_1ST_CLASS", "C0"}
		case "VLAK_23":
			seatClasses = []string{"TRAIN_R23_LOW_COST", "TRAIN_R23_STANDARD", "TRAIN_R23_RELAX", "TRAIN_R23_BUSINESS"}
		default:
			slog.Error("Unsupported line group code in YELLOW_R8", "lineGroupCode", lineGroupCode)
			c.JSON(http.StatusNotImplemented, gin.H{"error": "Currently unsupported line group code. Please contact support."})
			return
		}
	default:
		slog.Error("Unsupported vehicle standard key", "vehicleStandardKey", vehicleStandardKey)
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Currently unsupported vehicle standard key. Please contact support."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"seatClasses": seatClasses})
}
