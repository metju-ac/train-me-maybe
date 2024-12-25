package main

import (
	"fmt"
	"github.com/metju-ac/train-me-maybe/internal/handlers"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func main() {
	fmt.Println("Hello, RegioJet!")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	selectedRoutes, err := handlers.HandleRouteSelection(apiClient)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Selected routes:", selectedRoutes)
}
