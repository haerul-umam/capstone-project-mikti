package repository

import (
	"errors"

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

func (repo *OrderRepositoryImpl) GetOrder(Id string) (domain.Order, error) {
	var orderData domain.Order

	err := repo.db.First(&orderData, "id = ?", Id).Error

	if err != nil {
		return domain.Order{}, errors.New("order tidak ditemukan")
	}

	return orderData, nil
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

func (repo *OrderRepositoryImpl) GetDetailOrder(Id string) (domain.Category, domain.Order, error) {
	var orderData domain.Order

	err := repo.db.First(&orderData, "id = ?", Id).Error

	if err != nil {
		return domain.Category{}, domain.Order{}, errors.New("order tidak ditemukan")
	}

	errUser := repo.db.Model(&domain.Order{}).Preload("User").Find(&orderData).Error

	if errUser != nil {
		return domain.Category{}, domain.Order{}, errUser
	}

	var eventData domain.Event

	errEvent := repo.db.First(&eventData, "id = ?", orderData.EventID).Error

	if errEvent != nil {
		return domain.Category{}, domain.Order{}, errEvent
	}

	var categoryData domain.Category

	errCategory := repo.db.First(&categoryData, "id = ?", eventData.CategoryID).Error

	if errCategory != nil {
		return domain.Category{}, domain.Order{}, errEvent
	}

	return categoryData, orderData, nil
}

func (repo *OrderRepositoryImpl) ChangeOrderStatus(order domain.Order) (domain.Order, error) {

	err := repo.db.Model(domain.Order{}).Where("id = ?", order.OrderID).Updates(order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}
