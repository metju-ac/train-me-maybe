package cli

import (
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

var ErrRouteNotFound = errors.New("route not found")

func formatRoute(route *openapiclient.SimpleRoute) string {
	return fmt.Sprintf(
		"%s - %s (Transfers: %d)",
		route.DepartureTime.Format("02.01. 15:04"),
		route.ArrivalTime.Format("02.01. 15:04"),
		route.GetTransfersCount(),
	)
}

func SelectRoute(routes []openapiclient.SimpleRoute) (*openapiclient.SimpleRoute, error) {
	var selectedRoute string

	formattedRoutes := make([]string, 0, len(routes))
	for _, route := range routes {
		formattedRoutes = append(formattedRoutes, formatRoute(&route))
	}

	selectPrompt := &survey.Select{
		Message: "Select a route:",
		Options: formattedRoutes,
	}
	err := survey.AskOne(selectPrompt, &selectedRoute)
	if err != nil {
		return nil, fmt.Errorf("failed to ask for route selection: %w", err)
	}

	for _, route := range routes {
		formatted := formatRoute(&route)
		if formatted == selectedRoute {
			return &route, nil
		}
	}

	return nil, fmt.Errorf("%w", ErrRouteNotFound)
}
