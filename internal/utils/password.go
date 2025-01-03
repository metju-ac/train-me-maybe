package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

func GenerateSalt() (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %v", err)
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func HashPassword(password, salt string) string {
	saltBytes, _ := base64.StdEncoding.DecodeString(salt)
	hash := pbkdf2.Key([]byte(password), saltBytes, 4096, 32, sha256.New)
	return base64.StdEncoding.EncodeToString(hash)
}
