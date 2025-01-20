package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

const (
	Iterations = 4096
	KeyLength  = 32
)

func GenerateSalt() (string, error) {
	salt := make([]byte, KeyLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func HashPassword(password, salt string) string {
	saltBytes, _ := base64.StdEncoding.DecodeString(salt)
	hash := pbkdf2.Key([]byte(password), saltBytes, Iterations, KeyLength, sha256.New)
	return base64.StdEncoding.EncodeToString(hash)
}
