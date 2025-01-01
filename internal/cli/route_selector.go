package cli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func FormatRoute(route *openapiclient.SimpleRoute) string {
	return fmt.Sprintf("%s - %s (Transfers: %d)", route.DepartureTime.Format("02.01. 15:04"), route.ArrivalTime.Format("02.01. 15:04"), route.GetTransfersCount())
}

func SelectRoute(routes []openapiclient.SimpleRoute) (*openapiclient.SimpleRoute, error) {
	var selectedRoute string

	formattedRoutes := make([]string, 0, len(routes))
	for _, route := range routes {
		formattedRoutes = append(formattedRoutes, FormatRoute(&route))
	}

	selectPrompt := &survey.Select{
		Message: "Select a route:",
		Options: formattedRoutes,
	}
	err := survey.AskOne(selectPrompt, &selectedRoute)
	if err != nil {
		return nil, err
	}

	for _, route := range routes {
		formatted := FormatRoute(&route)
		if formatted == selectedRoute {
			return &route, nil
		}
	}

	return nil, fmt.Errorf("route not found")
}
