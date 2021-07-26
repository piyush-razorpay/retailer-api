package models

type Product struct {
	ID        string `json:"id" gorm:"primary_key"`
	Name      string `json:"naame" gorm:"type:varchar(50)"`
	Quantity  int    `json:"quantity" gorm:"type:varchar(50)"`
	Price     int    `json:"price" gorm:"type:int"`
	CreatedAt string `json:"createdat" gorm:"type:varchar(100)"`
}

type CreateProductParams struct {
	Name     string `json:"username"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
