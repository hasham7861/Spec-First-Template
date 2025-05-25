package utils

import (
	"encoding/json"
	"net/http"
)

func EncodeResponseToJSON(httpWriter http.ResponseWriter, data interface{}) {

	// this writes to the http response writer automatically
	err := json.NewEncoder(httpWriter).Encode(data)
	
	if err != nil {
		http.Error(httpWriter, err.Error(), http.StatusInternalServerError)
	}
}

func EncodeErrorToJSON(httpWriter http.ResponseWriter, errorMessage string, statusCode int) {
	httpWriter.WriteHeader(statusCode)

	response := map[string]string{"error": errorMessage}
	
	err := json.NewEncoder(httpWriter).Encode(response)
	if err != nil {
		http.Error(httpWriter, err.Error(), http.StatusInternalServerError)
	}
}