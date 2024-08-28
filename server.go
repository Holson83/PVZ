package main

import (
	"PVZ/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/orders", handlers.GetAllOrder)                  //1
	e.GET("/orders/:id", handlers.GetOrderId)               //1
	e.POST("/orders", handlers.CreateOrder)                 //1
	e.PUT("/orders/:id", handlers.UpdateOrder)              //1
	e.DELETE("/orders/:id", handlers.DeleteOrder)           //1
	e.PUT("/orders/:id/statuses", handlers.SetOrderStatus)  //1
	e.PUT("/orders/:id/products", handlers.AddProductOrder) //1

	e.GET("/statuses", handlers.GetAllStatus)        //1
	e.GET("/statuses/:id", handlers.GetStatusId)     //1
	e.POST("/statuses", handlers.CreateStatus)       //1
	e.PUT("/statuses/:id", handlers.UpdateStatus)    //1
	e.DELETE("/statuses/:id", handlers.DeleleStatus) //1

	e.GET("/products", handlers.GetAllProduct)        //1
	e.GET("/products/:id", handlers.GetProductId)     //1
	e.POST("/products", handlers.CreateProduct)       //1
	e.PUT("/products/:id", handlers.UpdateProduct)    //1
	e.DELETE("/products/:id", handlers.DeleleProduct) //1

	e.GET("/baskets/:id", handlers.GetBasketId)
	e.POST("/baskets", handlers.CreateBasket)
	e.PUT("/baskets/:id", handlers.UpdateBasket)
	e.DELETE("/baskets/:id", handlers.DeleleBasket)
	e.PUT("/baskets/:id/products", handlers.AddProductBasket)

	e.Logger.Fatal(e.Start(":8080"))
}
