package handlers

import (
	"PVZ/models"

	"net/http"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"

	"strconv"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Order{})
}

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}

func GetAllOrders(c echo.Context) error {
	var orders []models.Order

	db.Find(&orders)

	return c.JSON(http.StatusOK, orders)
}

func GetAllOrdersId(c echo.Context) error {
	var order models.Order

	db.Last(&order, c.Param("id"))

	return c.JSON(http.StatusOK, order)
}

func PostAllOrders(c echo.Context) error {
	fullName := c.FormValue("fullName")
	id := StringToUint(c.FormValue("id"))

	var existingRecord models.Order
	db.Take(&existingRecord, id)

	if id == existingRecord.ID {
		return c.String(http.StatusBadRequest, "Order not created")
	}

	order := models.Order{FullName: fullName, ID: id}
	db.Create(&order)
	return c.String(http.StatusCreated, "Order created")

	//return c.String(http.StatusInternalServerError, "Pizdec")
}

func DeleleOrder(c echo.Context) error {
	id := StringToUint(c.FormValue("id"))

	var record models.Order
	db.Take(&record, id)

	if id != record.ID {
		return c.String(http.StatusBadRequest, "Not found order")
	}

	order := models.Order{ID: id}
	db.Delete(&order)
	return c.String(http.StatusOK, "Delete Order")
}

func UpdateOrder(c echo.Context) error {
	fullName := c.FormValue("fullName")
	id := StringToUint(c.FormValue("id"))

	var urecord models.Order
	db.Take(&urecord, id)

	if id != urecord.ID {
		return c.String(http.StatusBadRequest, "Not found order")
	}

	order := models.Order{FullName: fullName, ID: id}

	db.Save(&order)

	return c.String(http.StatusOK, "Update Order")
}
