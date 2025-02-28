package config

import (
	"context"
	"log"

	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// Firebase struct to store authentication client
type Firebase struct {
	AuthDatabase *auth.Client
}

// Global Firebase instance
var FB Firebase

// InitFirebase initializes Firebase authentication
func InitFirebase() {
	opt := option.WithCredentialsFile("authwithfirebase-d23d5-firebase-adminsdk-fbsvc-1b2b588077.json") // Ensure this file exists
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error getting Firebase auth client: %v", err)
	}

	FB.AuthDatabase = authClient
}
