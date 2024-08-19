package main

import (
	"PVZ/handlers"

	//"net/http"

	"github.com/labstack/echo/v4"
	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"
	// "strconv"
)

func main() {
	//handlers.Init()

	e := echo.New()

	e.GET("/orders", handlers.GetAllOrders)

	e.GET("/orders/:id", handlers.GetAllOrdersId)

	e.POST("/orders", handlers.PostAllOrders)

	e.DELETE("/orders", handlers.DeleleOrder)

	e.PUT("/orders", handlers.UpdateOrder)

	e.Logger.Fatal(e.Start(":8080"))
}
