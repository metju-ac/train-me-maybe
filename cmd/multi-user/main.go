package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/database"
	"github.com/metju-ac/train-me-maybe/internal/handlers/http"
	"github.com/metju-ac/train-me-maybe/internal/middleware"
	"github.com/metju-ac/train-me-maybe/internal/repositories"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
	"os"
	"time"
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

	userRepo := repositories.NewUserRepository(db)
	handler := &http.Handler{UserRepo: userRepo}

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
	}

	publicRoutes := router.Group("/api")
	{
		publicRoutes.POST("/login", handler.Login)
		publicRoutes.POST("/register", handler.Register)
	}

	router.Run(":8080")
}
