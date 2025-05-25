GO ?= go
GOPATH ?= $(shell $(GO) env GOPATH)
TOOLS_BIN := $(GOPATH)/bin
OAPI_CODEGEN := $(TOOLS_BIN)/oapi-codegen
AIR := $(TOOLS_BIN)/air

# list all commands to not create object files for
.PHONY: tools generate clean start dev

# Install all tools defined in tools.go
tools:
		@echo "Installing tools from tools.go..."
		$(GO) mod download
		grep -E '^\t_' tools.go | cut -d'"' -f2 | xargs -t -I % $(GO) install %@latest

generate: tools
		@echo "Generating code from OpenAPI spec..."
		$(OAPI_CODEGEN) -generate std-http-server,spec,types -package gen -o api/api.gen.go api/api.yaml

clean:
		rm -f api/api.gen.go

start: 
		@echo "Starting the server..."
		$(GO) run cmd/main.go

dev: tools
		@echo "Starting development server with hot reload..."
		$(AIR)