package http

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type UpdateUserRequest struct {
	CutOffTime     *int    `json:"cutOffTime"`
	MinimalCredit  *int    `json:"minimalCredit"`
	CreditUser     *string `json:"creditUser"`
	CreditPassword *string `json:"creditPassword"`
	TariffClass    *string `json:"tariffKey"`
}

func (h *Handler) GetUser(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		slog.Warn("Unauthorized access attempt")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := h.UserRepo.FindByEmail(email.(string))
	if err != nil {
		slog.Error("Error retrieving user", "email", email, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		return
	}

	slog.Info("User retrieved successfully", "email", user.Email)
	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		slog.Warn("Unauthorized access attempt")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Invalid request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := h.UserRepo.FindByEmail(email.(string))
	if err != nil {
		slog.Error("Error retrieving user", "email", email, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		return
	}

	user.CutOffTime = req.CutOffTime
	user.MinimalCredit = req.MinimalCredit
	user.CreditUserNumber = req.CreditUser
	user.CreditUserPassword = req.CreditPassword
	user.DefaultTariffClass = req.TariffClass

	if err := h.UserRepo.Update(user); err != nil {
		slog.Error("Error updating user", "email", email, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	slog.Info("User updated successfully", "email", user.Email)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
