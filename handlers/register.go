package handlers

import (
	"net/http"
	"sample-spec-first-api/gen"
)

// RegisterHandlers maps all routes defined in the OpenAPI spec to their respective methods.
func RegisterHandlers(mux *http.ServeMux, h gen.ServerInterface) {
	// TODO: an improvement would be to auto register all routes by file name
	mux.HandleFunc("/", h.GetIndex)
	mux.HandleFunc("/hello", h.GetHello)

	// Add more routes here as your API expands
	// mux.HandleFunc("/something", h.PostSomething)
}
