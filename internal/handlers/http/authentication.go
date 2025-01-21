package http

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/metju-ac/train-me-maybe/internal/dbmodels"
	"github.com/metju-ac/train-me-maybe/internal/utils"
)

type UserCredentials struct {
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (h *Handler) Login(c *gin.Context) {
	var credentials UserCredentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		slog.Error("Invalid data", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	credentials.Email = strings.ToLower(credentials.Email)

	user, err := h.UserRepo.FindByEmail(credentials.Email)
	if err != nil {
		slog.Warn("Invalid credentials", "email", credentials.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	passwordHash := utils.HashPassword(credentials.Password, user.Salt)
	if passwordHash != user.PasswordHash {
		slog.Warn("Invalid credentials", "email", credentials.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken([]byte(h.Config.General.JwtSecretKey), user.Email)
	if err != nil {
		slog.Error("Error generating token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	slog.Info("User logged in successfully", "email", user.Email)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) Register(c *gin.Context) {
	var credentials UserCredentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		slog.Error("Invalid data", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	credentials.Email = strings.ToLower(credentials.Email)

	if _, err := h.UserRepo.FindByEmail(credentials.Email); err == nil {
		slog.Warn("User with given email already exists", "email", credentials.Email)
		c.JSON(http.StatusConflict, gin.H{"error": "User with given email already exists"})
		return
	}

	salt, err := utils.GenerateSalt()
	if err != nil {
		slog.Error("Error generating salt", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating salt"})
		return
	}

	passwordHash := utils.HashPassword(credentials.Password, salt)

	user := dbmodels.User{
		Email:        credentials.Email,
		PasswordHash: passwordHash,
		Salt:         salt,
	}

	if err := h.UserRepo.Create(&user); err != nil {
		slog.Error("Error saving user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user"})
		return
	}

	token, err := utils.GenerateToken([]byte(h.Config.General.JwtSecretKey), user.Email)
	if err != nil {
		slog.Error("Error generating token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	slog.Info("User registered successfully", "email", user.Email)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "token": token})
}
