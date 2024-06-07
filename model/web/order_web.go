package web

import (
	"github.com/haerul-umam/capstone-project-mikti/model/domain"
	"github.com/haerul-umam/capstone-project-mikti/model/entity"
)

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

type OrdersPageRequest struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type OrdersPageResponse struct {
	Total       int64                      `json:"total"`
	TotalPages  int                        `json:"totalPages"`
	CurrentPage int                        `json:"currentPage"`
	Orders      []entity.OrderOnPageEntity `json:"orders"`
}
