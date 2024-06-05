package web

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
