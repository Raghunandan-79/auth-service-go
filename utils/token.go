package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRefreshToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}