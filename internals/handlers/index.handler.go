package handlers

import (
	"net/http"
	"sample-spec-first-api/utils"
)


func (ar *ApiRegistry) GetIndex(response http.ResponseWriter, request *http.Request) {
	utils.EncodeResponseToJSON(response, IndexResponse{Status: "API is running"})
}