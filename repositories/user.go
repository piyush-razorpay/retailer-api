package repositories

import (
	"fmt"
	"github.com/Gandhi24/retailer-api/models"
	"github.com/jinzhu/gorm"
	"time"
)

func CreateUser(arg *models.CreateUserParams, db *gorm.DB) (models.User, error) {
	fmt.Println(arg)
	newUser := &models.User{
		Username:       arg.Username,
		HashedPassword: arg.HashedPassword,
		Email:          arg.Email,
		FullName:       arg.FullName,
		CreatedAt:      time.Now().String(),
	}
	result := db.Create(&newUser)
	return *newUser, result.Error
}

func GetUser(username string, db *gorm.DB) (models.User, error) {
	var user models.User
	result := db.Where("Username = ?", username).First(&user)
	return user, result.Error
}
