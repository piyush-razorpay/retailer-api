package controllers

import (
	"fmt"
	"github.com/Gandhi24/retailer-api/models"
	"github.com/Gandhi24/retailer-api/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createProductRequest struct {
	Name     string `json:"name" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
	Price    int    `json:"price" binding:"required"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := models.ProductParams{
		Name:     req.Name,
		Quantity: req.Quantity,
		Price:    req.Price,
	}

	fmt.Println(arg)
	product, err := repositories.CreateProduct(&arg, server.connection)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, newProductResponse(product))
}

func (server *Server) getAllProducts(ctx *gin.Context) {
	products, err := repositories.GetProducts(server.connection)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (server *Server) getProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := repositories.GetProductById(id, server.connection)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, product)
}

type updateProductRequest struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type productResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

func newProductResponse(product models.Product) productResponse {
	fmt.Println("product: ", product)
	return productResponse{
		Id:       product.ProductID,
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    product.Price,
	}
}

func (server *Server) updateProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var req updateProductRequest
	fmt.Println("id: ", id)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req)

	arg := models.ProductParams{
		Name:     req.Name,
		Quantity: req.Quantity,
		Price:    req.Price,
	}

	product, err := repositories.UpdateProductById(id, server.connection, &arg)
	fmt.Println("err: ", err)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newProductResponse(product)
	fmt.Println("rsp: ", rsp)
	ctx.JSON(http.StatusOK, rsp)
}
