package web

import "github.com/haerul-umam/capstone-project-mikti/model/domain"

type OrderRequest struct {
	EventID       int            `validate:"required" json:"event_id"`
	Quantity      int            `validate:"required" json:"quantity"`
	PaymentMethod domain.Payment `validate:"required" json:"payment_method"`
}

type OrderResponse struct {
	OrderID string        `json:"order_id"`
	Amount  int           `json:"amount"`
	Status  domain.Status `json:"status"`
}
