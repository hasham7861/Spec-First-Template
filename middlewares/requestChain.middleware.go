package middlewares

import (
	"net/http"
)

// A middleware that can take list of middlewares and apply to given handler one by one.
func chain(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func ApplyPresetMiddlewaresToHandler(handler http.HandlerFunc) http.HandlerFunc {
	return chain(handler, applyHeaders, applyLogging)
}