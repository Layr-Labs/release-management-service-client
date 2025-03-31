package client

import (
	"net/http"
	"time"
)

func NewClientConfig(endpointUrl string, environment string, duration time.Duration, client *http.Client) *Config {
	return &Config{
		endpointUrl,
		environment,
		duration,
		client,
	}
}

type Config struct {
	EndpointUrl string
	Environment string
	Timeout     time.Duration
	HTTPClient  *http.Client
}
