package models

type Product struct {
	ProductID string `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"type:varchar(50);not null"`
	Quantity  int    `json:"quantity" gorm:"type:varchar(50);not null"`
	Price     int    `json:"price" gorm:"type:int;not null"`
	CreatedAt string `json:"created_at" gorm:"type:varchar(100)"`
}

type ProductParams struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
