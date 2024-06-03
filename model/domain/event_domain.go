package domain

import "time"

type Event struct {
	Id          int    `gorm:"column:id;primaryKey:type:serial"`
	CategoryID  int    `gorm:"column:category_id"`
	Name        string `gorm:"column:name"`
	Date        string `gorm:"column:date"`
	Price       int    `gorm:"column:price"`
	IsFree      bool   `gorm:"column:is_free"`
	City        string `gorm:"city"`
	Description string `gorm:"description"`
	Quota       int    `gorm:"quota"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func (event *Event) TableName() string {
	return "event"
}
