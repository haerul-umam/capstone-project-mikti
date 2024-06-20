package web

import (
	"github.com/haerul-umam/capstone-project-mikti/model/entity"
	"gorm.io/gorm"
)

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

type Filter string

const (
	Termurah   Filter = "termurah"
	Termahal   Filter = "termahal"
	Terpopuler Filter = "terpopuler"
	Terbaru    Filter = "terbaru"
)

type AllEventDataRequest struct {
	PriceMax   int
	PriceMin   int
	City       string
	Date       string
	CategoryId int
	Filter     Filter
	Limit      int
	Page       int
}

type AllEventDataResponse struct {
	Total       int64                `json:"total"`
	TotalPages  int                  `json:"totalPages"`
	CurrentPage int                  `json:"currentPage"`
	Events      []entity.EventEntity `json:"events"`
}
