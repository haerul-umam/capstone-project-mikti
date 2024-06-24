package controller

import (
	"net/http"
	"strconv"

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
	limit, _ := strconv.Atoi(e.QueryParam("limit"))
	page, _ := strconv.Atoi(e.QueryParam("page"))

	queryParams := web.OrdersPageRequest{
		Page: page,
		Limit: limit,
	}

	if err := e.Validate(queryParams); err != nil {
		return err
	}

	orderData, err := controller.orderService.GetOrderListOnPage(queryParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return e.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "Sukses", orderData))
}

func (controller *OrderControllerImpl) DetailOrder(e echo.Context) error {
	getOrder, errGetOrder := controller.orderService.GetDetailOrder(e.Param("id"), helper.GetClaimsValue(e).ID)

	if errGetOrder != nil {
		return e.JSON(http.StatusNotFound, web.ResponseToClient(http.StatusNotFound, errGetOrder.Error(), nil))
	}

	return e.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "success", getOrder))
}

func (controller *OrderControllerImpl) ChangeOrderStatus(e echo.Context) error {
	stat := new(web.ChangePaymentRequest)

	if err := e.Bind(stat); err != nil {
		return e.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := e.Validate(stat); err != nil {
		return err
	}

	changeStatus, errChange := controller.orderService.ChangeOrderStatus(e.Param("id"), *stat)

	if errChange != nil {
		return e.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, errChange.Error(), nil))
	}

	return e.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "success", changeStatus))
}

func (controller *OrderControllerImpl) GetAllPayment(e echo.Context) error {
	limit, _ := strconv.Atoi(e.QueryParam("limit"))
	page, _ := strconv.Atoi(e.QueryParam("page"))
	status := e.QueryParam("status")

	queryParams := web.AllPaymentQueryRequest{
		Status: web.StatusPayment(status),
		Page: page,
		Limit: limit,
	}

	if err := e.Validate(queryParams); err != nil {
		return err
	}

	data, err := controller.orderService.GetAllPayment(queryParams)

	if err != nil {
		return e.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return e.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "sukses", data))
}