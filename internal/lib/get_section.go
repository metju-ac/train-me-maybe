package lib

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/metju-ac/train-me-maybe/internal/models"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func getSectionIdsFromRoute(routeIds string) ([]int64, error) {
	intIds := make([]int64, 0)

	ids := strings.Split(routeIds, ",")
	for _, idStr := range ids {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid route ID: %s", idStr)
		}
		intIds = append(intIds, id)
	}

	return intIds, nil
}

func GetSection(input *models.UserInput) (*openapiclient.SimpleSection, error) {

	sectionIds, err := getSectionIdsFromRoute(input.SelectedRouteIds)
	if err != nil {
		slog.Error("Error getting section IDs", "error", err)
		return nil, fmt.Errorf("error getting section IDs: %v", err)
	}

	// For now, we only support exactly one section of the given route
	if len(sectionIds) != 1 {
		slog.Error("Expected one section ID, got", "sectionIds", sectionIds)
		return nil, fmt.Errorf("expected one section ID, got %v", sectionIds)
	}

	return openapiclient.NewSimpleSection(sectionIds[0], input.DepartingStation.StationID, input.ArrivingStation.StationID), nil
}
