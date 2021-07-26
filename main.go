package main

import (
	"github.com/Gandhi24/retailer-api/controllers"
	models2 "github.com/Gandhi24/retailer-api/models"
	"github.com/Gandhi24/retailer-api/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type database struct {
	connection *gorm.DB
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func main() {
	db, err := gorm.Open(util.DBDriver, util.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	db.AutoMigrate(&models2.User{}, &models2.Product{})

	server, err := controllers.NewServer(db)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(util.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
