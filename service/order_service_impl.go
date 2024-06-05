package service

import (
	"errors"
	"math"

	"github.com/haerul-umam/capstone-project-mikti/model/domain"
	"github.com/haerul-umam/capstone-project-mikti/model/entity"
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
		EventID:       getEvent.EventID,
		UserID:        userID,
		NameEvent:     getEvent.Name,
		DateEvent:     getEvent.Date,
		PriceEvent:    getEvent.Price,
		IsFree:        getEvent.Is_free,
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

func (service *OrderServiceImpl) GetOrderListOnPage(request web.OrdersPageRequest) (web.OrdersPageResponse, error) {
	orderReq := web.OrdersPageRequest{
		Limit: request.Limit,
		Page:  request.Page,
	}
	getOrderList, total, errGetOrderList := service.repository.GetOrdersPage(orderReq.Limit, orderReq.Page)

	if errGetOrderList != nil {
		return web.OrdersPageResponse{}, errGetOrderList
	}

	totalPages := int(math.Ceil(float64(total) / float64(orderReq.Limit)))
	orders := entity.ToOrderListOnPageEntities(getOrderList)

	return web.OrdersPageResponse{
		Total:       total,
		TotalPages:  totalPages,
		CurrentPage: orderReq.Page,
		Orders:      orders,
	}, nil
}

func (service *OrderServiceImpl) GetDetailOrder(Id string, userID string) (web.DetailOrderResponse, error) {
	category, getOrder, errGetOrder := service.repository.GetDetailOrder(Id)

	if errGetOrder != nil {
		return web.DetailOrderResponse{}, errGetOrder
	}

	if userID != getOrder.UserID {
		return web.DetailOrderResponse{}, errors.New("user tidak sesuai")
	}

	return web.DetailOrderResponse{
		OrderID:       getOrder.OrderID,
		Costumer:      web.Costumer{Name: getOrder.User.Name},
		NameEvent:     getOrder.NameEvent,
		Quantity:      getOrder.Quantity,
		Amount:        getOrder.Amount,
		DateEvent:     getOrder.DateEvent,
		PriceEvent:    getOrder.PriceEvent,
		IsFree:        getOrder.IsFree,
		City:          getOrder.City,
		Description:   getOrder.Description,
		Category:      web.Category{Name: category.Name},
		StatusPayment: getOrder.Status,
		PaymentMethod: getOrder.PaymentMethod,
	}, nil
}
