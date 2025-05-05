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

### Folder Architecture

```
.
├── README.md          # Project documentation
├── Makefile           # Build and run scripts
├── go.mod             # Go module definition
├── go.sum             # Go module dependencies
├── tools.go           # Tooling dependencies for development
├── api
│   ├── spec.yaml      # OpenAPI specification
│   └── generated      # Generated code from OpenAPI spec
├── cmd
│   └── main.go        # Entry point for the application
├── internal
│   ├── handlers       # HTTP handlers
│   ├── services       # Business logic
│   └── models         # Data models
└── pkg
  └── utils          # Utility functions
```
