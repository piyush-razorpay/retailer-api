package controllers

import (
	"fmt"
	"github.com/Gandhi24/retailer-api/models"
	"github.com/Gandhi24/retailer-api/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type createOrderRequest struct {
	UserId    string `json:"user_id" binding:"required,alphanum"`
	ProductId string `json:"product_id" binding:"required,alphanum"`
	Quantity  int    `json:"quantity" binding:"required,numeric"`
}

type orderResponse struct {
	OrderId   string `json:"order_id"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Status    string `json:"order_status"`
}

func newOrderResponse(order models.Order, orderStatus string) orderResponse {
	return orderResponse{
		OrderId:   order.OrderID,
		Quantity:  order.Quantity,
		ProductId: order.ProductId,
		Status:    orderStatus,
	}
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req createOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := models.CreateOrderParams{
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	}

	err := repositories.ValidUser(req.UserId, server.connection)
	fmt.Println(err)
	if err != nil {
		ctx.Set("message", "No user with given userId")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": ctx.Keys["message"],
		})
		return
	}

	_, err = repositories.GetProductById(req.ProductId, server.connection)
	if err != nil {
		ctx.Set("message", "No product with given productId")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": ctx.Keys["message"],
		})
		return
	}

	lastOrderTime, err := repositories.GetLastOrderTime(req.UserId, server.connection)
	if err != nil {
		if err.Error() != "record not found" {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		lastOrderTime = time.Now().Add(-time.Hour * 100)
	}

	currentTime := time.Now()
	fmt.Print("current time is: ", currentTime)
	if currentTime.Before(lastOrderTime.Add(time.Minute * 5)) {
		ctx.Set("message", "Can't order just yet!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": ctx.Keys["message"],
		})
		return
	}

	fmt.Println(arg)
	order, err := repositories.CreateOrder(&arg, server.connection)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newOrderResponse(order, "order_placed")
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) getAllOrders(ctx *gin.Context) {
	orders, err := repositories.GetOrders(server.connection)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var responses []orderResponse
	for _, order := range orders {
		responses = append(responses, newOrderResponse(order, "order_placed"))
	}
	ctx.JSON(http.StatusOK, responses)
}

func (server *Server) getOrdersByUserID(ctx *gin.Context) {
	userId := ctx.Param("id")
	fmt.Println("userID: ", userId)
	err := repositories.ValidUser(userId, server.connection)
	fmt.Println(err)
	if err != nil {
		ctx.Set("message", "No user with provided userId")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": ctx.Keys["message"],
		})
		return
	}

	orders, err := repositories.GetOrdersByUserID(userId, server.connection)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var responses []orderResponse
	for _, order := range orders {
		responses = append(responses, newOrderResponse(order, "order_placed"))
	}
	ctx.JSON(http.StatusOK, responses)
}
