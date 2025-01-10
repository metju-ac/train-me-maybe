package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/database"
	"github.com/metju-ac/train-me-maybe/internal/handlers"
	"github.com/metju-ac/train-me-maybe/internal/handlers/http"
	"github.com/metju-ac/train-me-maybe/internal/middleware"
	"github.com/metju-ac/train-me-maybe/internal/models"
	"github.com/metju-ac/train-me-maybe/internal/repositories"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
	"os"
	"time"
)

func fetchStations(apiClient *openapiclient.APIClient) []models.StationModel {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stations, err := handlers.FetchStations(ctx, apiClient)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return stations
}

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

	userRepo := repositories.NewUserRepository(db)
	watchedRouteRepo := repositories.NewWatchedRouteRepository(db)

	handler := &http.Handler{
		Stations:          make(map[int64]models.StationModel),
		UserRepo:          userRepo,
		WatchedRouteRepo:  watchedRouteRepo,
		ApiClient:         apiClient,
		GoroutineContexts: make(map[uint]context.CancelFunc),
		Config:            *config,
	}

	stations := fetchStations(apiClient)
	for _, station := range stations {
		handler.Stations[station.StationID] = station
	}

	err = handler.RestartWatchedRoutes()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Update with your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Protected routes (require authentication)
	protectedRoutes := router.Group("/api/auth")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		protectedRoutes.GET("/user", handler.GetUser)
		protectedRoutes.PUT("/user", handler.UpdateUser)

		protectedRoutes.GET("/station", handler.GetStations)
		protectedRoutes.GET("/route", handler.GetRoutes)

		protectedRoutes.POST("/watchedRoute", handler.CreateWatchedRoute)
		protectedRoutes.GET("/watchedRoute", handler.GetWatchedRoutes)
		protectedRoutes.DELETE("/watchedRoute/:id", handler.DeleteWatchedRoute)
	}

	publicRoutes := router.Group("/api")
	{
		publicRoutes.POST("/login", handler.Login)
		publicRoutes.POST("/register", handler.Register)
	}

	router.Run(":8080")
}
