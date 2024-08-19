package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID uint
	//Status   uint
	FullName string
	//ID_product uint
}
