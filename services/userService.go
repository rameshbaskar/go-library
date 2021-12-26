package services

import (
	"fmt"
	"go-library/database"
	"go-library/models"
	"go-library/utils"
)

type UserService struct {

}

func (us UserService) FindOrCreateUser(
	username string,
	plainTextPassword string,
	fullName string,
	email string,
) (models.User, error) {
	existingUser, err := us.FindByUsername(username)
	if err != nil {
		fmt.Println(err)
	}

	if existingUser.Username == username {
		fmt.Println("User already exists. No need to create again.")
		return existingUser, nil
	}

	fmt.Println("User does not exist. Need to create a new one.")
	newUser := models.User{
		Username: username,
		EncPassword: utils.EncryptPassword(plainTextPassword),
		FullName: fullName,
		Email: email,
		CreatedBy: username,
		UpdatedBy: username,
	}
	db := database.GetDbConnection()
	createResult := db.Omit("id").Create(&newUser)
	db.Commit()
	return newUser, createResult.Error
}

func (us UserService) FindByUsername(username string) (models.User, error) {
	db := database.GetDbConnection()
	var user models.User
	result := db.Model(models.User{}).Find(&user, "username = ?", username)
	return user, result.Error
}
