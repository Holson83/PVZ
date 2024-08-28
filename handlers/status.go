package handlers

import (
	"PVZ/models"

	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllStatus(c echo.Context) error {
	var status []models.Status

	db.Find(&status)

	return c.JSON(http.StatusOK, status)
}

func GetStatusId(c echo.Context) error {
	var status models.Status

	db.Take(&status, c.Param("id"))

	if status.ID == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, status)
}

func CreateStatus(c echo.Context) error {
	status := models.Status{Type: c.FormValue("type")}

	db.Create(&status)

	return c.String(http.StatusCreated, "Status created")
}

func UpdateStatus(c echo.Context) error {
	var status models.Status

	db.Take(&status, c.Param("id"))

	if status.ID == 0 {
		return c.String(http.StatusBadRequest, "Not found status")
	}

	status.Type = c.FormValue("type")

	db.Save(&status)

	return c.JSON(http.StatusOK, status)
}

func DeleleStatus(c echo.Context) error {
	var status models.Status

	db.Take(&status, c.Param("id"))

	if status.ID == 0 {
		return c.String(http.StatusBadRequest, "Not found status")
	}

	db.Delete(&status)

	return c.String(http.StatusOK, "Delete status")
}
