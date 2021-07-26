package controllers

import (
	"database/sql"
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

	arg := models.CreateProductParams{
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
	ctx.JSON(http.StatusOK, product)
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
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, product)
}
