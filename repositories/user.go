package repositories

import (
	"fmt"
	"github.com/Gandhi24/retailer-api/models"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
	"time"
)

func CreateUser(arg *models.CreateUserParams, db *gorm.DB) (models.User, error) {
	fmt.Println(arg)
	newUser := &models.User{
		UserID:         generateUserId(),
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
	result := db.Where("username = ?", username).First(&user)
	return user, result.Error
}

func ValidUser(userId string, db *gorm.DB) error {
	var user models.User
	result := db.First(&user, "user_id = ?", userId)
	return result.Error
}

func generateUserId() string {
	guid := xid.New()
	return "USER" + guid.String()
}
