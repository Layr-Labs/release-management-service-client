package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Layr-Labs/release-management-service-client/pkg/gen"
	"github.com/Layr-Labs/release-management-service-client/pkg/model"
)

type ReleaseManagementServiceClient interface {
	ListAvsReleaseKeys(ctx context.Context, avsId string) (*model.ListAvsReleaseKeysResponse, error)
	ListOperatorReleases(ctx context.Context, operatorId string) (*model.ListOperatorRequirementsResponse, error)
}

type Client struct {
	api *gen.ClientWithResponses
}

func NewClient(cfg *Config) (ReleaseManagementServiceClient, error) {
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: cfg.Timeout}
	}
	endpointUrl := cfg.EndpointUrl
	if endpointUrl == "" {
		env := cfg.Environment
		if env == "" {
			return nil, fmt.Errorf("no client, url or environment provided")
		}
		endpointUrl = getEndpointFromEnvironment(cfg.Environment)
	}

	client, err := gen.NewClientWithResponses(endpointUrl, gen.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}

	return &Client{api: client}, nil
}

func (c *Client) ListAvsReleaseKeys(ctx context.Context, avsId string) (*model.ListAvsReleaseKeysResponse, error) {
	resp, err := c.api.ListAvsReleaseKeysWithResponse(ctx, avsId)
	if err != nil {
		return nil, fmt.Errorf("release API request failed: %w", err)
	}

	if resp == nil {
		return nil, fmt.Errorf("release API returned a nil response")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("release API returned status %d: %s", resp.StatusCode(), string(resp.Body))
	}

	if resp.JSON200 == nil {
		return nil, fmt.Errorf("release API returned no JSON body")
	}

	if resp.JSON200.AvsReleasePublicKeys == nil {
		return &model.ListAvsReleaseKeysResponse{Keys: []string{}}, nil
	}

	return &model.ListAvsReleaseKeysResponse{
		Keys: *resp.JSON200.AvsReleasePublicKeys,
	}, nil
}

func (c *Client) ListOperatorReleases(ctx context.Context, operatorId string) (*model.ListOperatorRequirementsResponse, error) {
	resp, err := c.api.ListOperatorReleasesWithResponse(ctx, operatorId)
	if err != nil {
		return nil, fmt.Errorf("release API request failed: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("release API returned status %d: %s", resp.StatusCode(), string(resp.Body))
	}

	if resp.JSON200 == nil {
		return nil, fmt.Errorf("release API returned no data (JSON200 was nil)")
	}

	if resp.JSON200.OperatorRequirements == nil {
		return nil, fmt.Errorf("release API returned an empty response body")
	}

	var result []model.OperatorApplication
	for _, req := range *resp.JSON200.OperatorRequirements {
		var components []model.Component
		if req.Components != nil {
			for _, component := range *req.Components {
				components = append(components, model.Component{
					Name:             safeStr(component.Name),
					Description:      safeStr(component.Description),
					Location:         safeStr(component.Location),
					LatestArtifactId: safeStr(component.LatestArtifactId),
					ReleaseTimestamp: safeStr(component.ReleaseTimestamp),
				})
			}
		}

		result = append(result, model.OperatorApplication{
			ApplicationName: safeStr(req.ApplicationName),
			OperatorSetId:   safeStr(req.OperatorSetId),
			Description:     safeStr(req.Description),
			Components:      components,
		})
	}

	return &model.ListOperatorRequirementsResponse{
		OperatorRequirements: result,
	}, nil
}

func safeStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func getEndpointFromEnvironment(environment string) string {
	switch environment {
	case "prod":
		return "https://api.eigenlayer.xyz/release-management-service"
	case "preprod":
		return "https://api.preprod.eigenlayer.xyz/release-management-service"
	case "testnet":
		return "https://api.testnet.eigenlayer.xyz/release-management-service"
	default:
		return "http://localhost:8080/release-management-service"
	}
}
