package main

import (
	"context"
	"fmt"

	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func main() {
	fmt.Println("Hello, RegioJet!")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	_, httpRes, _ := apiClient.ConstsAPI.GetActionPrices(context.Background()).Execute()

	fmt.Println("GOT", httpRes.StatusCode)
}
