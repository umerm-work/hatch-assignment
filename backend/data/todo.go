package data

import (
	"gorm.io/gorm"
)

type Todo struct {
	ID          int64  `gorm:"id"`
	Title       string `gorm:"title,size:50"`
	Description string `gorm:"description,size:250"`
	Priority    int    `gorm:"priority"`
	gorm.Model
}
