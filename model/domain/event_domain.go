package domain

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	EventID     int `gorm:"column:id;primaryKey;autoIncrement"`
	CategoryID  int
	Category    Category `gorm:"foreignKey:category_id"`
	Name        string
	Date        string
	Price       int
	Is_free     bool
	City        string
	Description string
	Quota       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (event *Event) TableName() string {
	return "event"
}
