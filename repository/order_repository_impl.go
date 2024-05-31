package repository

import (
	"github.com/haerul-umam/capstone-project-mikti/model/domain"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db}
}

func (repo *OrderRepositoryImpl) CreateOrder(order domain.Order) (domain.Order, error) {
	err := repo.db.Create(&order).Error

	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
