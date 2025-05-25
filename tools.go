//go:build tools

// tools.go keeps track of build tools as dependencies
package tools

import (
	_ "github.com/air-verse/air"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
)
