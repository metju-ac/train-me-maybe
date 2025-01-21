package utils

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt"
)

const hoursInDay = 24

var (
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrInvalidToken         = errors.New("invalid token")
)

func GenerateToken(secretKey []byte, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * hoursInDay).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		slog.Error("Error signing token", "error", err)
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	slog.Info("Token generated successfully", "email", email)
	return signedToken, nil
}

func VerifyToken(secretKey []byte, tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w", ErrInvalidSigningMethod)
		}
		return secretKey, nil
	})
	if err != nil {
		slog.Error("Error parsing token", "error", err)
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		slog.Info("Token verified successfully", "email", claims["email"])
		return claims, nil
	}

	slog.Warn("Invalid token")
	return nil, fmt.Errorf("%w", ErrInvalidToken)
}
