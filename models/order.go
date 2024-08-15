package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID       uint
	FullName string
}
