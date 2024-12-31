package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashUsernamePassword(username, password string) string {
	combined := username + ":" + password
	hash := sha256.New()
	hash.Write([]byte(combined))
	hashBytes := hash.Sum(nil)
	hashHex := hex.EncodeToString(hashBytes)

	return hashHex
}

func GetCurrentRegion() (string, error) {
	return "Canada", nil
}
