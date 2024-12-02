package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func validateEnvVars() error {
	required := []string{
		"FIREBASE_PROJECT_ID",
		"FIREBASE_PRIVATE_KEY_ID",
		"FIREBASE_PRIVATE_KEY",
		"FIREBASE_CLIENT_EMAIL",
		"FIREBASE_CLIENT_ID",
	}

	for _, v := range required {
		if os.Getenv(v) == "" {
			return fmt.Errorf("missing required environment variable: %s", v)
		}
	}
	return nil
}

func InitFirebase() (*firebase.App, error) {
	if err := validateEnvVars(); err != nil {
		return nil, err
	}

	ctx := context.Background()

	bucketName := os.Getenv("FIREBASE_STORAGE_BUCKET")
	config := &firebase.Config{
		StorageBucket: bucketName,
	}

	// Create service account JSON from environment variables
	serviceAccount := map[string]string{
		"type":                        "service_account",
		"project_id":                  os.Getenv("FIREBASE_PROJECT_ID"),
		"private_key_id":              os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		"private_key":                 os.Getenv("FIREBASE_PRIVATE_KEY"),
		"client_email":                os.Getenv("FIREBASE_CLIENT_EMAIL"),
		"client_id":                   os.Getenv("FIREBASE_CLIENT_ID"),
		"auth_uri":                    "https://accounts.google.com/o/oauth2/auth",
		"token_uri":                   "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":        fmt.Sprintf("https://www.googleapis.com/robot/v1/metadata/x509/%s", os.Getenv("FIREBASE_CLIENT_EMAIL")),
	}

	// Convert to JSON
	serviceAccountJSON, err := json.Marshal(serviceAccount)
	if err != nil {
		log.Fatalf("error creating service account JSON: %v\n", err)
		return nil, err
	}

	// Create credentials option
	opt := option.WithCredentialsJSON(serviceAccountJSON)

	// Initialize Firebase app
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return nil, err
	}

	return app, nil
}
