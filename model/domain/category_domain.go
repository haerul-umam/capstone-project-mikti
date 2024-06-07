package domain

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `gorm:"column:name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (category *Category) TableName() string {
	return "category_event"
}
