package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log/slog"
	"time"
)

// TODO: env var
var secretKey = []byte("secretpassword")

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		slog.Error("Error signing token", "error", err)
		return "", err
	}

	slog.Info("Token generated successfully", "email", email)
	return signedToken, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		slog.Error("Error parsing token", "error", err)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		slog.Info("Token verified successfully", "email", claims["email"])
		return claims, nil
	}

	slog.Warn("Invalid token")
	return nil, fmt.Errorf("Invalid token")
}
