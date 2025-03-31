package client

import (
	"fmt"
	"net/http"

	"github.com/Layr-Labs/release-management-service-client/pkg/gen"
)

type Client struct {
	api *gen.ClientWithResponses
}

func NewClient(cfg Config) (*Client, error) {
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: cfg.Timeout}
	}
	baseURL := cfg.BaseURL
	if baseURL == "" {
		env := cfg.Environment
		if env == "" {
			return nil, fmt.Errorf("no client, url or environment provided")
		}
		baseURL = getUrlFromEnvironment(cfg.Environment)
	}

	client, err := gen.NewClientWithResponses(cfg.BaseURL, gen.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}

	return &Client{api: client}, nil
}

func getUrlFromEnvironment(environment string) string {
	switch environment {
	case "prod":
		return "https://api.eigenlayer.xyz/release-management-service/"
	case "preprod":
		return "https://api.preprod.eigenlayer.xyz/release-management-service/"
	case "testnet":
		return "https://api.testnet.eigenlayer.xyz/release-management-service/"
	default:
		return "http://localhost:2345/release-management-service/"
	}
}
