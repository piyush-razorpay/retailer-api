package repositories

import (
	"fmt"
	"github.com/Gandhi24/retailer-api/models"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
	"time"
)

func CreateProduct(arg *models.ProductParams, db *gorm.DB) (models.Product, error) {
	fmt.Println(arg)
	newProduct := &models.Product{
		ProductID: generateProductId(),
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
	result := db.First(&product, "product_id = ?", id)
	return product, result.Error
}

func UpdateProductById(id string, db *gorm.DB, arg *models.ProductParams) (models.Product, error) {
	var product models.Product
	fmt.Println(id)
	result := db.First(&product, "product_id = ?", id)
	updatedName := product.Name
	updatedPrice := product.Price
	updatedQuantity := product.Quantity
	if arg.Name != "" {
		updatedName = arg.Name
	}
	if arg.Price != 0 {
		updatedPrice = arg.Price
	}
	if arg.Quantity != 0 {
		updatedQuantity = arg.Quantity
	}
	db.Model(&product).Updates(models.Product{Name: updatedName, Price: updatedPrice, Quantity: updatedQuantity})

	return product, result.Error
}

func generateProductId() string {
	guid := xid.New()
	return "PROD" + guid.String()
}
