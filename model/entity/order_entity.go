package entity

type OrderEntity struct {
	OrderID     string `json:"order_id"`
	EventID     int    `json:"event_id"`
	NameEvent   string `json:"name_event"`
	DateEvent   string `json:"date_event"`
	PriceEvent  string `json:"price_event"`
	IsFree      bool   `json:"is_free"`
	Description string `json:"description"`
	City        string `json:"city"`
	Quantity    int    `json:"quantity"`
}
