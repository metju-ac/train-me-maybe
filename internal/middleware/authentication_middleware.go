package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/metju-ac/train-me-maybe/internal/utils"
	"log/slog"
	"net/http"
	"strings"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			slog.Warn("Missing authentication token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})
			c.Abort()
			return
		}

		// The token should be prefixed with "Bearer "
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			slog.Warn("Invalid authentication token format", "token", tokenString)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
			c.Abort()
			return
		}

		tokenString = tokenParts[1]

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			slog.Warn("Invalid authentication token", "error", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
			c.Abort()
			return
		}

		slog.Info("Authentication token verified", "email", claims["email"])
		c.Set("email", claims["email"])
		c.Next()
	}
}
