package services

import (
	"TODO/databases"
	"TODO/models"
	"errors"
)

func GetUserByUsernameAndPassword(username string, password string) (*models.User, error) {
	user := &models.User{}

	err := databases.DB.Where("username = ? AND password = ?", username, password).First(user).Error
	if err != nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}
