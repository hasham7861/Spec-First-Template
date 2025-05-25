package handlers

import (
	"log"
	"net/http"
	"sample-spec-first-api/db/repositories"
	"sample-spec-first-api/utils"
	"strconv"
)

func (ar *ApiRegistry) GetUser(response http.ResponseWriter, request *http.Request) {
	userID := request.URL.Query().Get("id")
	
	id, err := strconv.Atoi(userID)
	if err != nil {
		log.Println("Error converting user ID to integer:", err)
		utils.EncodeErrorToJSON(response, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := repositories.GetUserByID(uint(id))
	if err != nil {
		utils.EncodeErrorToJSON(response, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}
	
	utils.EncodeResponseToJSON(response, user)
}