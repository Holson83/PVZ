package models

import (
	"gorm.io/gorm"
)

type ProductBasket struct {
	gorm.Model
	BasketID  uint    `gorm:"primaryKey"`
	ProductID uint    `gorm:"primaryKey"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Count     uint    `gorm:"default:1"`
}
