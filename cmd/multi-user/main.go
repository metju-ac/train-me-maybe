package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

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
)

const (
	maxAge         = 12 * time.Hour
	contextTimeout = 30 * time.Second
)

func fetchStations(apiClient *openapiclient.APIClient, languages []string) map[string][]models.StationModel {
	stationsMap := make(map[string][]models.StationModel)

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)

	for _, lang := range languages {
		stations, err := handlers.FetchStations(ctx, apiClient, lang)
		if err != nil {
			slog.Error("Error fetching stations for language", "language", lang, "error", err)
			cancel()
			os.Exit(1)
		}

		stationsMap[lang] = stations
	}

	cancel()
	return stationsMap
}

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		slog.Error("Error loading configuration", "error", err)
		os.Exit(1)
	}

	db, err := database.ConnectAndMigrate(config.DB)
	if err != nil {
		slog.Error("Error connecting to database and applying migrations", "error", err)
		os.Exit(1)
	}

	slog.Info("Connected to database and migrations applied", "database", db)

	apiConfiguration := openapiclient.NewConfiguration(config.General.APIBaseURL)
	apiClient := openapiclient.NewAPIClient(apiConfiguration)

	userRepo := repositories.UserRepository{DB: db, Key: []byte(config.General.EncryptionKey)}
	watchedRouteRepo := repositories.WatchedRouteRepository{DB: db}
	successfullPurchaseRepo := repositories.SuccessfulPurchaseRepository{DB: db}
	notificationRepo := repositories.NotificationRepository{DB: db}

	stations := fetchStations(apiClient, config.General.Languages)

	handler := &http.Handler{
		Stations:               make(map[int64]models.StationModel),
		LangStations:           stations,
		UserRepo:               userRepo,
		WatchedRouteRepo:       watchedRouteRepo,
		SuccessfulPurchaseRepo: successfullPurchaseRepo,
		NotificationRepo:       notificationRepo,
		APIClient:              apiClient,
		GoroutineContexts:      make(map[uint]context.CancelFunc),
		Config:                 *config,
	}

	for _, stationList := range stations {
		for _, station := range stationList {
			handler.Stations[station.StationID] = station
		}
		break
	}

	err = handler.RestartWatchedRoutes()
	if err != nil {
		slog.Error("Error restarting watched routes", "error", err)
		os.Exit(1)
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.General.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           maxAge,
	}))

	// Protected routes (require authentication)
	protectedRoutes := router.Group("/api/auth")
	protectedRoutes.Use(middleware.AuthenticationMiddleware([]byte(config.General.JwtSecretKey)))
	{
		protectedRoutes.GET("/user", handler.GetUser)
		protectedRoutes.PUT("/user", handler.UpdateUser)

		protectedRoutes.GET("/station", handler.GetStations)
		protectedRoutes.GET("/route", handler.GetRoutes)

		protectedRoutes.POST("/watchedRoute", handler.CreateWatchedRoute)
		protectedRoutes.GET("/watchedRoute", handler.GetWatchedRoutes)
		protectedRoutes.DELETE("/watchedRoute/:id", handler.DeleteWatchedRoute)

		protectedRoutes.GET("/seatClass", handler.GetSeatClasses)
	}

	publicRoutes := router.Group("/api")
	{
		publicRoutes.POST("/login", handler.Login)
		publicRoutes.POST("/register", handler.Register)
	}

	err = router.Run(fmt.Sprintf(":%d", config.General.ServerPort))
	if err != nil {
		slog.Error("Error starting the server", "error", err)
		os.Exit(1)
	}
}
