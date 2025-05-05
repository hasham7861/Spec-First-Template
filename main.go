package main

import (
	"log"
	"net/http"

	"sample-spec-first-api/handlers"
)

func main() {
	handler := handlers.NewAPIHandler()

	mux := http.NewServeMux()
	handlers.RegisterHandlers(mux, handler)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
