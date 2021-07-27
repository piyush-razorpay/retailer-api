package controllers

import (
	"fmt"
	"github.com/Gandhi24/retailer-api/token"
	"github.com/Gandhi24/retailer-api/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	connection *gorm.DB
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(db *gorm.DB) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(util.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		connection: db,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.POST("/products", server.createProduct)
	router.GET("/products", server.getAllProducts)
	router.GET("/products/:id", server.getProductByID)

	router.POST("/orders", server.createOrder)
	//router.GET("/orders", server.getAllOrders)
	//router.GET("/orders/:id", server.getOrderByID)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
