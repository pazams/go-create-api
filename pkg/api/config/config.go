package config

import (
	"context"
	"fmt"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// Config holds the application configuration
type Config struct {
	AppPort          string
	AppEnv           string
	APIToken         string
	PgConnectionName string
	PgAddr           string
	PgUser           string
	PgPassword       string
	PgDatabase       string
}

// New returns a new Config object
func New() (*Config, error) {
	var accessor func(string) string
	var err error

	if os.Getenv("IS_GCP_CONFIG") != "" {
		fmt.Println("using gcp config")
		projectID := os.Getenv("GCP_PROJECT")
		accessor, err = createGCPSecretAccessor(projectID) // TODO error if now proj Id
		if err != nil {
			return nil, err
		}
	} else {
		fmt.Println("using env config")
		accessor = os.Getenv
	}

	return &Config{
		// PORT var name is mendated by GAE https://cloud.google.com/appengine/docs/standard/go112/runtime
		AppPort:          get("PORT", "", os.Getenv),
		AppEnv:           get("APP_ENV", "", accessor),
		APIToken:         get("API_TOKEN", "", accessor),
		PgConnectionName: get("POSTGRES_GCP_CONNECTION_NAME", "", accessor), // GCP cloud SQL format "project:zone:instance"
		PgAddr:           get("POSTGRES_ADDR", "default", accessor),         // for integration tests
		PgUser:           get("POSTGRES_USER", "", accessor),
		PgPassword:       get("POSTGRES_PASSWORD", "", accessor),
		PgDatabase:       get("POSTGRES_DATABASE", "", accessor),
	}, nil
}

func get(key, defaultValue string, accessor func(string) string) string {
	value := accessor(key)
	if value == "" {
		if defaultValue == "" {
			panic(fmt.Sprintf("Config for %v: No value or default value", key))
		}
		return defaultValue
	}
	return value
}

func createGCPSecretAccessor(projectID string) (func(string) string, error) {
	if projectID == "" {
		return nil, fmt.Errorf("No projectIDfound for accessing secrets")
	}

	return func(key string) string {
		value, err := accessGCPSecret(projectID, key)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return value
	}, nil
}

func accessGCPSecret(projectID, key string) (string, error) {
	name := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, key)

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secretmanager client: %v", err)
	}

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %v", err)
	}

	return string(result.Payload.Data), nil
}
