package domain

import "time"

type Category struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `gorm:"column:name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (category *Category) TableName() string {
	return "category_event"
}
