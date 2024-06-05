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

func (repo *OrderRepositoryImpl) GetOrdersPage(limit int, page int) ([]domain.Order, int64, error) {
	var orders []domain.Order
	offset := (page - 1) * limit

	errData := repo.db.Limit(limit).Offset(offset).Find(&orders).Error

	if errData != nil {
		return []domain.Order{}, 0, errData
	}

	var total int64
	errTotal := repo.db.Model(&domain.Order{}).Count(&total).Error
	if errTotal != nil {
		return []domain.Order{}, 0, errTotal
	}

	return orders, total, nil
}
