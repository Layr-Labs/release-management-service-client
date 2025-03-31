OPENAPI_SPEC ?= internal/api/release-management-service-api.yaml

generate-api:
	@echo "Generating OpenAPI client and types from $(OPENAPI_SPEC)..."
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	@mkdir -p pkg/gen
	@oapi-codegen -package gen -generate types $(OPENAPI_SPEC) > pkg/gen/types.gen.go
	@oapi-codegen -package gen -generate client $(OPENAPI_SPEC) > pkg/gen/client.gen.go
	@go fmt pkg/gen/*.gen.go