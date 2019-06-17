package config

import (
	"fmt"

	"github.com/spf13/viper"
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
func New() *Config {
	viper.AutomaticEnv()
	return &Config{
		AppPort:          get("port", ""), // PORT var name is mendated by GAE https://cloud.google.com/appengine/docs/standard/go112/runtime
		AppEnv:           get("app_env", ""),
		APIToken:         get("api_token", ""),
		PgConnectionName: get("postgres_gcp_connection_name", ""), // GCP cloud SQL format "project:zone:instance"
		PgAddr:           get("postgres_addr", "default"),         // for integration tests
		PgUser:           get("postgres_user", ""),
		PgPassword:       get("postgres_password", ""),
		PgDatabase:       get("postgres_database", ""),
	}
}

func get(key, defaultValue string) string {
	value := viper.GetString(key)
	if value == "" {
		if defaultValue == "" {
			panic(fmt.Sprintf("Config for %v: No value or default value", key))
		}
		return defaultValue
	}
	return value
}
