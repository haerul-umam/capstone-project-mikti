package service

import "github.com/haerul-umam/capstone-project-mikti/model/web"

type OrderService interface {
	CreateOrder(request web.OrderRequest, userID string) (web.OrderResponse, error)
	GetOrderListOnPage(request web.OrdersPageRequest) (web.OrdersPageResponse, error)
	GetDetailOrder(Id string, userID string) (web.DetailOrderResponse, error)
	ChangeOrderStatus(Id string, status web.ChangePaymentRequest) (map[string]interface{}, error)
	GetAllPayment(request web.AllPaymentQueryRequest) (web.AllPaymentDataResponse, error)
}
