package repositories

import (
	"fmt"
	"github.com/Gandhi24/retailer-api/models"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
	"time"
)

func CreateOrder(arg *models.CreateOrderParams, db *gorm.DB) (models.Order, error) {
	fmt.Println(arg)
	newOrder := &models.Order{
		OrderID:   generateOrderId(),
		ProductId: arg.ProductId,
		UserId:    arg.UserId,
		Quantity:  arg.Quantity,
		CreatedAt: time.Now().String(),
	}
	result := db.Create(&newOrder)
	return *newOrder, result.Error
}

func generateOrderId() string {
	guid := xid.New()
	return "ORD" + guid.String()
}
