package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sendydwi/online-book-store/api"
	"github.com/sendydwi/online-book-store/config/database"
	"github.com/sendydwi/online-book-store/services/cart"
	"github.com/sendydwi/online-book-store/services/order"
	"github.com/sendydwi/online-book-store/services/product"
	"github.com/sendydwi/online-book-store/services/user"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	r := gin.Default()
	db := database.Init()

	RegisterService(r, db)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8182")
}

func RegisterService(r *gin.Engine, db *gorm.DB) {
	userHandler := user.NewRestHandler(db)
	productHandler := product.NewRestHandler(db)
	cartHandler := cart.NewRestHandler(db)
	orderHandler := order.NewRestHandler(db)

	handlers := []api.Handler{userHandler, productHandler, cartHandler, orderHandler}

	for _, handler := range handlers {
		handler.RegisterHandler(r)
	}
}
