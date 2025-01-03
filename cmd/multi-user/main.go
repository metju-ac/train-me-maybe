package main

import (
	"fmt"
	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/database"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
	"os"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	db, err := database.ConnectAndMigrate()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Connected to database and migrations applied", db)
	apiConfiguration := openapiclient.NewConfiguration(config.General.ApiBaseUrl)
	apiClient := openapiclient.NewAPIClient(apiConfiguration)

	fmt.Println("Hello, multi-user!")
	fmt.Println(apiClient)
}
