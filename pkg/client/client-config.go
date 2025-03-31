package client

import (
	"net/http"
	"time"
)

func NewClientConfig(baseUrl string, environment string, duration time.Duration, client *http.Client) *Config {
	return &Config{
		baseUrl,
		environment,
		duration,
		client,
	}
}

type Config struct {
	BaseURL     string
	Environment string
	Timeout     time.Duration
	HTTPClient  *http.Client
}
