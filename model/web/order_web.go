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
	Limit      int						`query:"limit" validate:"required,gte=1,lte=100"`
	Page       int 						`query:"page" validate:"required,gte=1"`
}

type OrdersPageResponse struct {
	Total       int64                      `json:"total"`
	TotalPages  int                        `json:"totalPages"`
	CurrentPage int                        `json:"currentPage"`
	Orders      []entity.OrderOnPageEntity `json:"orders"`
}

type Costumer struct {
	Name string `json:"name"`
}

type DetailOrderResponse struct {
	OrderID       string         `json:"order_id"`
	Costumer      Costumer       `json:"costumer"`
	NameEvent     string         `json:"name_event"`
	Quantity      int            `json:"quantity"`
	Amount        int            `json:"amount"`
	DateEvent     string         `json:"date_event"`
	PriceEvent    int            `json:"price_event"`
	IsFree        bool           `json:"is_free"`
	City          string         `json:"city"`
	Description   string         `json:"description"`
	Category      Category       `json:"category"`
	StatusPayment domain.Status  `json:"status_payment"`
	PaymentMethod domain.Payment `json:"payment_method"`
}

type ChangePaymentRequest struct {
	Status string `validate:"required,status_check" json:"status"`
}

type StatusPayment string

const (
	Menunggu   StatusPayment = "MENUNGGU"
	Diterima   StatusPayment = "DITERIMA"
	Ditolak 	 StatusPayment = "DITOLAK"
)

type AllPaymentQueryRequest struct {
	Status     StatusPayment	`query:"status" validate:"required,status_check"`
	Limit      int						`query:"limit" validate:"required,gte=1,lte=100"`
	Page       int 						`query:"page" validate:"required,gte=1"`
}

type AllPayment struct {
	OrderID       string         `json:"order_id"`
	Costumer      Costumer       `json:"costumer"`
	NameEvent     string         `json:"name_event"`
	DateEvent     string         `json:"date_event"`
	CreatedAt			string				 `json:"order_date"`
	Amount        int            `json:"amount"`
	City          string         `json:"city"`
	StatusPayment domain.Status  `json:"status_payment"`
}

type AllPaymentDataResponse struct {
	Total       int64                `json:"total"`
	TotalPages  int                  `json:"totalPages"`
	CurrentPage int                  `json:"currentPage"`
	Payments    []AllPayment 				 `json:"payments"`
}

func ToPaymentList(payments []domain.Order) []AllPayment {
	data := []AllPayment{}

	for _, payment := range payments {
		data = append(data, AllPayment{
			OrderID: payment.OrderID,
			Costumer: Costumer{payment.User.Name},
			NameEvent: payment.NameEvent,
			DateEvent: payment.DateEvent,
			CreatedAt: payment.CreatedAt.String(),
			Amount: payment.Amount,
			City: payment.City,
			StatusPayment: payment.Status,
		})
	}

	return data
}