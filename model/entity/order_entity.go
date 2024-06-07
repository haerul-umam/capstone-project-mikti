package entity

import "github.com/haerul-umam/capstone-project-mikti/model/domain"

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

type OrderOnPageEntity struct {
	OrderID   string        `json:"order_id"`
	NameEvent string        `json:"name_event"`
	Quantity  int           `json:"quantity"`
	Amount    int           `json:"amount"`
	DateEvent string        `json:"date_event"`
	Status    domain.Status `json:"status"`
}

func ToOrderOnPageEntity(order_id, name_event string, quantity, amount int, date_event string, status domain.Status) OrderOnPageEntity {
	return OrderOnPageEntity{order_id, name_event, quantity, amount, date_event, status}
}

func ToOrderListOnPageEntities(orders []domain.Order) []OrderOnPageEntity {
	data := []OrderOnPageEntity{}

	for _, order := range orders {
		data = append(data, ToOrderOnPageEntity(order.OrderID, order.NameEvent, order.Quantity, order.Amount, order.DateEvent, order.Status))
	}

	return data
}
