package handlers

import (
	"PVZ/models"

	"net/http"

	"github.com/labstack/echo/v4"

	"PVZ/utilities"
)

func GetAllProduct(c echo.Context) error {
	var product []models.Product

	db.Find(&product)

	return c.JSON(http.StatusOK, product)
}

func GetProductId(c echo.Context) error {
	var product models.Product

	db.Take(&product, c.Param("id"))

	if product.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	product := models.Product{
		Name:  c.FormValue("name"),
		Priсe: utilities.StringToUint(c.FormValue("price")),
	}

	db.Create(&product)

	return c.String(http.StatusCreated, "Product created")
}

func UpdateProduct(c echo.Context) error {
	var product models.Product

	db.Take(&product, c.Param("id"), c.Param("priсe"))

	if product.ID == 0 {
		return c.String(http.StatusBadRequest, "Not found product")
	}

	product.Name = c.FormValue("name")
	product.Priсe = utilities.StringToUint(c.FormValue("priсe"))

	db.Save(&product)

	return c.JSON(http.StatusOK, product)
}

func DeleleProduct(c echo.Context) error {
	var product models.Product

	db.Take(&product, c.Param("id"))

	if product.ID == 0 {
		return c.String(http.StatusBadRequest, "Not found product")
	}

	db.Delete(&product)

	return c.String(http.StatusOK, "Delete product")
}
