package handlers

import (
	"net/http"
	gen "sample-spec-first-api/api"
	"sample-spec-first-api/middlewares"
)

// RegisterHandlers maps all routes defined in the OpenAPI spec to their respective methods.
func RegisterHandlers(muxRouter *http.ServeMux, routeHandlers gen.ServerInterface) {
	muxRouter.HandleFunc("/", middlewares.ApplyPresetMiddlewaresToHandler(routeHandlers.GetIndex))
	muxRouter.HandleFunc("/user", middlewares.ApplyPresetMiddlewaresToHandler(routeHandlers.GetUser))
	// Add more routes here as your API expands
	// mux.HandleFunc("/something", h.PostSomething)
}
