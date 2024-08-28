package models

import (
	"gorm.io/gorm"
)

type Basket struct {
	gorm.Model
	Product []Product `gorm:"many2many:basket_product;"`
}
