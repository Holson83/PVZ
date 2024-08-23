package handlers

import (
	"PVZ/models"

	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllOrder(c echo.Context) error {
	var orders []models.Order

	db.Find(&orders)

	return c.JSON(http.StatusOK, orders)
}

func GetOrderId(c echo.Context) error {
	var order models.Order

	db.Preload("Status").Take(&order, c.Param("id"))

	if order.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, order)
}

func CreateOrder(c echo.Context) error {
	order := models.Order{FullName: c.FormValue("fullName")}

	db.Create(&order)

	return c.String(http.StatusCreated, "Order created")
}

func UpdateOrder(c echo.Context) error {
	var order models.Order

	db.Take(&order, c.Param("id"))

	if order.ID == 0 {
		return c.String(http.StatusBadRequest, "Not found order")
	}

	order.FullName = c.FormValue("fullName")

	db.Save(&order)

	return c.JSON(http.StatusOK, order)
}

func DeleleOrder(c echo.Context) error {
	var order models.Order

	db.Take(&order, c.Param("id"))

	if order.ID == 0 {
		return c.String(http.StatusBadRequest, "Not found order")
	}

	db.Delete(&order)

	return c.String(http.StatusOK, "Delete Order")
}

func SetOrderStatus(c echo.Context) error {
	var (
		order  models.Order
		status models.Status
	)

	db.Take(&order, c.Param("id"))
	db.Take(&status, c.FormValue("type"))

	if order.ID == 0 {
		return c.String(http.StatusNotFound, "Order not found")
	}

	if status.ID == 0 {
		return c.String(http.StatusBadRequest, "Status not found")
	}

	order.Status = status

	db.Save(&order)

	return c.JSON(http.StatusOK, order)
}
