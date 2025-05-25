# Spec-First-Template

## Description

- To use this project as a template for my golang apis

### Tech stack

- Go lang with std http/net with oapi codegen for auto routing and documentation

### Status

- **Status:** Not yet ready for development / production - so only for personal use only.

### Technical Notes

- `go mod tidy` to install all the tooling
- tools.go is where I keep dependencies for local development environment
- Makefile is used to create scripting to make development easier
- To run program `make start`

### Available Makefile Commands

| Command         | Description                                                             |
| --------------- | ----------------------------------------------------------------------- |
| `make tools`    | Install all tools defined in tools.go (oapi-codegen, air, etc.)         |
| `make generate` | Generate code from OpenAPI spec using oapi-codegen (creates api.gen.go) |
| `make clean`    | Remove generated code files (api.gen.go)                                |
| `make start`    | Start the server directly with `go run cmd/main.go`                     |
| `make dev`      | Start development server with hot reload using Air for faster iteration |

### Folder Architecture

```
.
├── README.md          # Project documentation
├── Makefile           # Build and run scripts
├── go.mod             # Go module definition
├── go.sum             # Go module dependencies
├── tools.go           # Tooling dependencies for development
├── sample-sqlite.db   # SQLite database file
├── api
│   ├── api.yaml       # OpenAPI specification
│   └── api.gen.go     # Generated code from OpenAPI spec
├── cmd
│   └── main.go        # Entry point for the application
├── db
│   ├── setup.go       # Database initialization
│   ├── models         # Database models
│   │   └── user.model.go
│   └── repositories   # Data access repositories
│       └── user.repository.go
├── internals          # Note: "internals" not "internal"
│   └── handlers       # HTTP handlers
│       ├── handler.types.go
│       ├── index.handler.go
│       ├── register.go # Handler registration functions
│       └── user.handler.go
├── middlewares        # HTTP middleware components
│   ├── logger.middleware.go
│   ├── requestChain.middleware.go
│   └── responseHeaders.middleware.go
└── utils              # Utility functions
    └── jsonEncoder.util.go
```
