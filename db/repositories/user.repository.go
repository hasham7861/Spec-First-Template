package repositories

import (
	"sample-spec-first-api/db"
	"sample-spec-first-api/db/models"
)


func CreateUser(user *models.User) error {
	result := db.GetDBInstance().Create(user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := db.GetDBInstance().First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}


func DeleteUser(id uint) error {
	return db.GetDBInstance().Delete(&models.User{}, id).Error
}