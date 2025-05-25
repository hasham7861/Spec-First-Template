package main

import (
	"log"
	"net/http"
	"sample-spec-first-api/db"
	"sample-spec-first-api/internals/handlers"
)


func main() {	
	mux := http.NewServeMux()

	apiRegistry := &handlers.ApiRegistry{}
	handlers.RegisterHandlers(mux, apiRegistry)

	db.InitDB("sample-sqlite.db")

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
