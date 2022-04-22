package repository

import (
	"crud_v2/app/database"
	"crud_v2/entity"
)

func FindUserById(id string, value ...string) *entity.User {
	var user entity.User
	if value != nil {
		database.Connector.Model(&user).Select(value).Where("id = ?", id).First(&user)
	} else {
		database.Connector.Model(&user).Where("id = ?", id).First(&user)
	}
	return &user
}

func FindUserByIdToSession(id string, value ...string) *entity.User {
	var user entity.User
	if value != nil {
		database.Connector.Model(&user).Select(value).Where("id = ?", id).First(&user)
	} else {
		database.Connector.Model(&user).Where("id = ?", id).First(&user)
	}
	return &user
}

func ValidUser(id string) bool {
	var user entity.User
	database.Connector.Model(&user).Select("id").Where("id = ?", id).First(&user)
	if user.Id == 0 {
		return false
	}
	return true
}
