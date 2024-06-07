package controller

import (
	"net/http"

	"github.com/haerul-umam/capstone-project-mikti/helper"
	"github.com/haerul-umam/capstone-project-mikti/model/web"
	"github.com/haerul-umam/capstone-project-mikti/service"
	"github.com/labstack/echo/v4"
)

type OrderControllerImpl struct {
	orderService service.OrderService
}

func NewOrderController(service service.OrderService, token helper.TokenUseCase) *OrderControllerImpl {
	return &OrderControllerImpl{
		orderService: service,
	}
}

func (controller *OrderControllerImpl) CreateOrder(e echo.Context) error {
	order := new(web.OrderRequest)

	if err := e.Bind(&order); err != nil {
		return e.JSON(400, web.ResponseToClient(400, err.Error(), nil))
	}

	if err := e.Validate(order); err != nil {
		return err
	}

	created_order, err := controller.orderService.CreateOrder(*order, helper.GetClaimsValue(e).ID)

	if err != nil {
		return e.JSON(
			400, web.ResponseToClient(400, err.Error(), nil),
		)
	}

	return e.JSON(200, web.ResponseToClient(200, "Sukses", created_order))
}

func (controller *OrderControllerImpl) GetOrdersPage(e echo.Context) error {
	order := new(web.OrdersPageRequest)

	if err := e.Bind(&order); err != nil {
		return e.JSON(400, web.ResponseToClient(400, err.Error(), nil))
	}

	if err := e.Validate(order); err != nil {
		return err
	}

	orderData, err := controller.orderService.GetOrderListOnPage(*order)
	if err != nil {
		return e.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return e.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "Sukses", orderData))
}
