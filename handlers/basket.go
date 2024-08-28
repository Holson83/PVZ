package handlers

import (
	"PVZ/models"

	"PVZ/utilities"

	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBasketId(c echo.Context) error {
	var basket models.Basket

	db.Preload("Product").Take(&basket, c.Param("id"))

	if basket.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, basket)
}

func CreateBasket(c echo.Context) error {
	basket := models.Basket{}

	db.Create(&basket)

	return c.JSON(http.StatusCreated, &basket)
}

func UpdateBasket(c echo.Context) error {
	var basket models.ProductBasket

	db.Take(&basket, c.Param("id"))

	if basket.ID == 0 {
		return c.String(http.StatusBadRequest, "Not found basket")
	}

	if basket.Count <= 0 {
		return c.String(http.StatusBadRequest, "Count must be greater than zero")
	}

	basket.Count = utilities.StringToUint(c.FormValue("count"))

	db.Save(&basket)

	return c.JSON(http.StatusOK, basket)
}

func DeleleBasket(c echo.Context) error {
	var basket models.Basket

	db.Take(&basket, c.Param("id"))

	if basket.ID == 0 {
		return c.String(http.StatusBadRequest, "Not found basket")
	}

	db.Delete(&basket)

	return c.String(http.StatusOK, "Delete basket")
}

func AddProductBasket(c echo.Context) error {
	var productbasket models.ProductBasket

	db.Where(&models.ProductBasket{
		BasketID:  utilities.StringToUint(c.Param("id")),
		ProductID: utilities.StringToUint(c.FormValue("productId")),
	}).Take(&productbasket)

	if productbasket.ProductID != 0 {
		return c.JSON(http.StatusBadRequest, productbasket.Count+1)
	}

	productbasket.ProductID = utilities.StringToUint(c.FormValue("productId"))
	productbasket.BasketID = utilities.StringToUint(c.Param("id"))

	db.Save(&productbasket)

	return c.JSON(http.StatusOK, productbasket)
}
