package repositories

import (
	"fmt"
	"github.com/Gandhi24/retailer-api/models"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
	"time"
)

func CreateProduct(arg *models.CreateProductParams, db *gorm.DB) (models.Product, error) {
	fmt.Println(arg)
	newProduct := &models.Product{
		ID:        generateProductId(),
		Name:      arg.Name,
		Quantity:  arg.Quantity,
		Price:     arg.Price,
		CreatedAt: time.Now().String(),
	}
	result := db.Create(&newProduct)
	return *newProduct, result.Error
}

func GetProducts(db *gorm.DB) ([]models.Product, error) {
	var products []models.Product
	result := db.Find(&products)
	return products, result.Error
}

func GetProductById(id string, db *gorm.DB) (models.Product, error) {
	var product models.Product
	result := db.First(&product, "id = ?", id)
	return product, result.Error
}

func generateProductId() string {
	guid := xid.New()
	return "PROD" + guid.String()
}
