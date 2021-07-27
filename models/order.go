package models

type Order struct {
	OrderID   string `json:"id" gorm:"primary_key"`
	UserId    string `json:"user_id" gorm:"not null"`
	ProductId string `json:"product_id" gorm:"not null"`
	Quantity  int    `json:"quantity" gorm:"not null"`
	CreatedAt string `json:"created_at" gorm:"type:varchar(100)"`
}

type CreateOrderParams struct {
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
