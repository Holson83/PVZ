package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name   string
	Priсe  uint
	Orders []Order `gorm:"many2many:order_product;"`
}
