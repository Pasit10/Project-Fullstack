package utils

import (
	"fmt"
	"log"

	"github.com/alexedwards/argon2id"
)

// HashPassword generates a hashed password using argon2id
func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return hash, nil
}

// VerifyPassword checks if the given password matches the hashed password
func VerifyPassword(password, hashedPassword string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hashedPassword)
	if err != nil {
		log.Printf("error verifying password: %v", err)
		return false, err
	}
	return match, nil
}
