package models

type User struct {
	ID             uint64 `json:"id" gorm:"primary_key;auto_increment""`
	Username       string `json:"usernaame" gorm:"type:varchar(50)"`
	FullName       string `json:"fullname" gorm:"type:varchar(50)"`
	HashedPassword string `json:"hashedpassword" gorm:"type:varchar(1000)"`
	CreatedAt      string `json:"createdat" gorm:"type:varchar(100)"`
	Email          string `json:"email" gorm:"type:varchar(50)"`
}

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}
