GO ?= go
GOPATH ?= $(shell $(GO) env GOPATH)
TOOLS_BIN := $(GOPATH)/bin
OAPI_CODEGEN := $(TOOLS_BIN)/oapi-codegen

.PHONY: tools generate clean start

# Install all tools defined in tools.go
tools:
		@echo "Installing tools from tools.go..."
		$(GO) mod download
		$(GO) install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

generate: tools
		@echo "Generating code from OpenAPI spec..."
		$(OAPI_CODEGEN) -generate std-http-server,spec,types -package gen -o gen/api.gen.go api.yaml

clean:
		rm -f gen/api.gen.go

start: 
		@echo "Starting the server..."
		$(GO) run main.go