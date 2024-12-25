package cli

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
	"strconv"
	"strings"
)

func formatRoute(route openapiclient.SimpleRoute) string {
	return fmt.Sprintf("%s - %s (Transfers: %d)", route.DepartureTime.Format("02.01. 15:04"), route.ArrivalTime.Format("02.01. 15:04"), route.GetTransfersCount())
}

func SelectRoute(routes []openapiclient.SimpleRoute) ([]int64, error) {
	var selectedRoute string

	formattedRoutes := make([]string, 0, len(routes))
	for _, route := range routes {
		formattedRoutes = append(formattedRoutes, formatRoute(route))
	}

	selectPrompt := &survey.Select{
		Message: "Select a route:",
		Options: formattedRoutes,
	}
	err := survey.AskOne(selectPrompt, &selectedRoute)
	if err != nil {
		return nil, err
	}

	var selectedIDs []int64
	for _, route := range routes {
		formatted := formatRoute(route)
		if formatted != selectedRoute {
			continue
		}

		ids := strings.Split(route.Id, ",")
		for _, idStr := range ids {
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid route ID: %s", idStr)
			}
			selectedIDs = append(selectedIDs, id)
		}
		break
	}

	if len(selectedIDs) == 0 {
		return nil, fmt.Errorf("route not found")
	}

	return selectedIDs, nil
}
