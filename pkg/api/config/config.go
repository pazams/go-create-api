package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	AppPort  string // any port number
	AppEnv   string // "test" or other string
	HasProxy bool   // "true" | "false" true should be set if deployed behind a reverse proxy. default to "false"
}

// New returns a new Config object
func New() (*Config, error) {
	accessor := os.Getenv

	hasProxy, err := strconv.ParseBool(get("HAS_PROXY", "false", accessor))
	if err != nil {
		return nil, err
	}

	return &Config{
		AppPort:  get("PORT", "", accessor),
		AppEnv:   get("APP_ENV", "", accessor),
		HasProxy: hasProxy,
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
