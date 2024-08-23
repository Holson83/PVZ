package main

import (
	"PVZ/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/orders", handlers.GetAllOrder)
	e.GET("/orders/:id", handlers.GetOrderId)
	e.POST("/orders", handlers.CreateOrder)
	e.DELETE("/orders", handlers.DeleleOrder)
	e.PUT("/orders/:id", handlers.UpdateOrder)

	e.GET("/products", handlers.GetAllProduct)
	e.GET("/products/:id", handlers.GetProductId)
	e.POST("/products", handlers.CreateProduct)
	e.DELETE("/products", handlers.DeleleProduct)
	e.PUT("/products/:id", handlers.UpdateProduct)

	e.GET("/statuses", handlers.GetAllStatus)
	e.GET("/statuses/:id", handlers.GetStatusId)
	e.POST("/statuses", handlers.CreateStatus)
	e.DELETE("/statuses", handlers.DeleleStatus)
	e.PUT("/statuses/:id", handlers.UpdateStatus)

	e.Logger.Fatal(e.Start(":8080"))
}
