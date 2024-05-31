package service

import (
	"errors"

	"github.com/haerul-umam/capstone-project-mikti/model/domain"
	"github.com/haerul-umam/capstone-project-mikti/model/web"
	"github.com/haerul-umam/capstone-project-mikti/repository"
)

type OrderServiceImpl struct {
	repository repository.OrderRepository
	event      repository.EventRepository
}

func NewOrderService(repository repository.OrderRepository, event repository.EventRepository) *OrderServiceImpl {
	return &OrderServiceImpl{
		repository: repository,
		event:      event,
	}
}

func (service *OrderServiceImpl) CreateOrder(request web.OrderRequest, userID string) (web.OrderResponse, error) {
	orderReq := web.OrderRequest{
		EventID:       request.EventID,
		Quantity:      request.Quantity,
		PaymentMethod: request.PaymentMethod,
	}
	getEvent, errGetEvent := service.event.GetEvent(orderReq.EventID)

	if errGetEvent != nil {
		return web.OrderResponse{}, errGetEvent
	}

	// check the qouta
	if getEvent.Quota < request.Quantity {
		return web.OrderResponse{}, errors.New("qouta tidak mencukupi")
	}

	// create order to db
	order := domain.Order{
		EventID:       getEvent.Id,
		UserID:        userID,
		NameEvent:     getEvent.Name,
		DateEvent:     getEvent.Date,
		PriceEvent:    getEvent.Price,
		IsFree:        getEvent.IsFree,
		Description:   getEvent.Description,
		City:          getEvent.City,
		Quantity:      orderReq.Quantity,
		PaymentMethod: orderReq.PaymentMethod,
		Amount:        orderReq.Quantity * getEvent.Price,
		Status:        "MENUNGGU",
	}

	createdOrder, errCreatedOrder := service.repository.CreateOrder(order)

	if errCreatedOrder != nil {
		return web.OrderResponse{}, errCreatedOrder
	}
	// decrease ticket qouta based on order quantity
	getEvent.Quota -= order.Quantity

	_, decreaseErr := service.event.DecreaseQouta(getEvent)

	if decreaseErr != nil {
		return web.OrderResponse{}, decreaseErr
	}

	return web.OrderResponse{
		OrderID: createdOrder.OrderID,
		Amount:  createdOrder.Amount,
		Status:  createdOrder.Status,
	}, nil
}
