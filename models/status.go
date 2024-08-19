package models

import (
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	Type string
}
