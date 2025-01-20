package lib

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/metju-ac/train-me-maybe/internal/models"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

var (
	ErrInvalidRouteID        = errors.New("invalid route ID")
	ErrInvalidSectionIDCount = errors.New("expected one section ID")
)

func getSectionIDsFromRoute(routeIDs string) ([]int64, error) {
	intIDs := make([]int64, 0)

	ids := strings.Split(routeIDs, ",")
	for _, idStr := range ids {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrInvalidRouteID, idStr)
		}
		intIDs = append(intIDs, id)
	}

	return intIDs, nil
}

func GetSection(input *models.UserInput) (*openapiclient.SimpleSection, error) {
	sectionIDs, err := getSectionIDsFromRoute(input.SelectedRouteIDs)
	if err != nil {
		slog.Error("Error getting section IDs", "error", err)
		return nil, fmt.Errorf("error getting section IDs: %w", err)
	}

	// For now, we only support exactly one section of the given route
	if len(sectionIDs) != 1 {
		slog.Error("Expected one section ID, got", "sectionIDs", sectionIDs)
		return nil, fmt.Errorf("%w: %v", ErrInvalidSectionIDCount, sectionIDs)
	}

	return openapiclient.NewSimpleSection(sectionIDs[0], input.DepartingStation.StationID, input.ArrivingStation.StationID), nil
}
