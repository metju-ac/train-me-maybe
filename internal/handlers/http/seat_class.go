package http

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const contextTimeout = 30 * time.Second

func (h *Handler) GetSeatClasses(c *gin.Context) {
	fromStationIDStr := c.Query("fromStationId")
	toStationIDStr := c.Query("toStationId")
	routeIDStr := c.Query("routeId")

	fromStationID, err := strconv.ParseInt(fromStationIDStr, 10, 64)
	if err != nil {
		slog.Error("Invalid fromStationId", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fromStationId"})
		return
	}

	toStationID, err := strconv.ParseInt(toStationIDStr, 10, 64)
	if err != nil {
		slog.Error("Invalid toStationID", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid toStationID"})
		return
	}

	slog.Info("Fetching route detail", "fromStationId", fromStationID, "toStationID", toStationID, "routeId", routeIDStr)
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	routeDetail, httpRes, err := h.APIClient.RoutesAPI.GetSimpleRouteDetail(ctx, routeIDStr).
		FromStationId(fromStationID).
		ToStationId(toStationID).
		Execute()
	if err != nil {
		slog.Error("Failed to fetch route detail", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching route detail"})
		return
	}
	defer httpRes.Body.Close()
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
