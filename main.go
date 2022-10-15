package main

import (
	"assigment2/helper"
	"assigment2/models/item"
	"assigment2/models/order"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := helper.DbConnect()
	helper.ErrorHelper(err)

	itemRepository := item.NewRepository(db)
	itemService := item.NewService(itemRepository)
	itemController := item.NewItemController(itemService)

	orderRepository := order.NewRepository(db)
	orderService := order.NewService(orderRepository)
	orderController := order.NewOrderController(orderService)

	router := gin.Default()

	api := router.Group("/api/")

	api.GET("/order", orderController.GetAll)
	api.POST("/order", orderController.Create)
	api.PUT("/order", orderController.Update)
	api.DELETE("/order/:id", orderController.Delete)

	api.GET("/item", itemController.Get)
	api.POST("/item", itemController.Create)

	router.Run(":8000")
}
