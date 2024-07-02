package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func generateToken(id int) string {
	// Get the current timestamp
	timestamp := time.Now().Unix()

	// Concatenate the user id and timestamp
	data := fmt.Sprintf("%d:%d", id, timestamp)

	// Create a SHA-256 hash of the concatenated data
	hash := sha256.New()
	hash.Write([]byte(data))

	// Convert the hash to a hexadecimal string
	token := hex.EncodeToString(hash.Sum(nil))

	return token
}
