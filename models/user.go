package models

type User struct {
	UserID         string `json:"id" gorm:"primary key"`
	Username       string `json:"username" gorm:"type:varchar(50);not null"`
	FullName       string `json:"full_name" gorm:"type:varchar(50)"`
	HashedPassword string `json:"hashed_password" gorm:"type:varchar(1000);not null"`
	CreatedAt      string `json:"created_at" gorm:"type:varchar(100)"`
	Email          string `json:"email" gorm:"type:varchar(50);not null"`
}

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}
