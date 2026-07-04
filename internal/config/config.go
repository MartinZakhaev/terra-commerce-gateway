package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Environment      string
	ListenAddress    string
	ShutdownTimeout  time.Duration
	CommerceBaseURL  string
	RequestTimeout   time.Duration
}

func Load() (Config, error) {
	shutdownTimeout, err := durationFromMilliseconds("TC_GATEWAY_SHUTDOWN_TIMEOUT_MS", 30_000)
	if err != nil {
		return Config{}, err
	}

	requestTimeout, err := durationFromMilliseconds("TC_GATEWAY_REQUEST_TIMEOUT_MS", 15_000)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		Environment:     valueOrDefault("TC_ENVIRONMENT", "local"),
		ListenAddress:   valueOrDefault("TC_GATEWAY_LISTEN_ADDR", ":8080"),
		ShutdownTimeout: shutdownTimeout,
		CommerceBaseURL: valueOrDefault("TC_COMMERCE_INTERNAL_URL", "http://localhost:9000"),
		RequestTimeout:  requestTimeout,
	}

	if cfg.ListenAddress == "" {
		return Config{}, fmt.Errorf("TC_GATEWAY_LISTEN_ADDR must not be empty")
	}

	return cfg, nil
}

func valueOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func durationFromMilliseconds(key string, fallback int64) (time.Duration, error) {
	raw := os.Getenv(key)
	if raw == "" {
		return time.Duration(fallback) * time.Millisecond, nil
	}

	value, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || value <= 0 {
		return 0, fmt.Errorf("%s must be a positive integer in milliseconds", key)
	}

	return time.Duration(value) * time.Millisecond, nil
}
