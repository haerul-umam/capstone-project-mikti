package web

import "gorm.io/gorm"

type EventUpdateServiceRequest struct {
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	Price       int    `json:"price"`
	Is_free     bool   `json:"is_free"`
	City        string `json:"city"`
	Description string `json:"description"`
	Quota       int    `json:"quota"`
}

type EventUpdateResponse struct {
	ItemID      int    `json:"item_id"`
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	Price       int    `json:"price"`
	Is_free     bool   `json:"is_free"`
	City        string `json:"city"`
	Description string `json:"description"`
	Quota       int    `json:"quota"`
}

type EventDetailResponse struct {
	ItemID      int      `json:"item_id"`
	CategoryID  int      `json:"category_id"`
	Name        string   `json:"name"`
	Date        string   `json:"date"`
	Price       int      `json:"price"`
	Is_free     bool     `json:"is_free"`
	City        string   `json:"city"`
	Description string   `json:"description"`
	Quota       int      `json:"quota"`
	Category    Category `json:"category"`
}

type EventDetailResponseAdmin struct {
	ItemID      int            `json:"item_id"`
	CategoryID  int            `json:"category_id"`
	Name        string         `json:"name"`
	Date        string         `json:"date"`
	Price       int            `json:"price"`
	Is_free     bool           `json:"is_free"`
	City        string         `json:"city"`
	Description string         `json:"description"`
	Quota       int            `json:"quota"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Category    Category       `json:"category"`
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
