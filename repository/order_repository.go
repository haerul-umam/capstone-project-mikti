package repository

import "github.com/haerul-umam/capstone-project-mikti/model/domain"

type OrderRepository interface {
	CreateOrder(order domain.Order) (domain.Order, error)
	GetOrder(Id string) (domain.Order, error)
	GetOrdersPage(limit int, page int) ([]domain.Order, int64, error)
	GetDetailOrder(Id string) (domain.Category, domain.Order, error)
	ChangeOrderStatus(order domain.Order) (domain.Order, error)
}
