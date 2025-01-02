package main

import (
	"fmt"
	"github.com/metju-ac/train-me-maybe/internal/config"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
	"os"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	apiConfiguration := openapiclient.NewConfiguration(config.General.ApiBaseUrl)
	apiClient := openapiclient.NewAPIClient(apiConfiguration)

	fmt.Println("Hello, multi-user!")
	fmt.Println(apiClient)
}
