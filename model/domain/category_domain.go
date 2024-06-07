package domain

import "time"

type Category struct {
	Id        int    `gorm:"column:id;primaryKey;type:serial"`
	Name      string `gorm:"column:name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (category *Category) TableName() string {
	return "category_event"
}
