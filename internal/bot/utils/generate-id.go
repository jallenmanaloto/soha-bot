package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateId() (string, error) {
	var id string
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		return id, err
	}

	return hex.EncodeToString(bytes), nil
}
