package controller

import (
	"net/http"
	"strconv"

	"github.com/haerul-umam/capstone-project-mikti/helper"
	"github.com/haerul-umam/capstone-project-mikti/model/web"
	"github.com/haerul-umam/capstone-project-mikti/service"
	"github.com/labstack/echo/v4"
)

type EventControllerImpl struct {
	eventService service.EventService
}

func NewEventController(service service.EventService) *EventControllerImpl {
	return &EventControllerImpl{
		eventService: service,
	}
}

func (controller *EventControllerImpl) CreateEvent(c echo.Context) error {
	eventBody := new(web.EventCreateServiceRequest)

	if err := c.Bind(eventBody); err != nil {
		return c.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(eventBody); err != nil {
		return err
	}

	event, err := controller.eventService.CreateEvent(*eventBody)

	if err != nil {
		return c.JSON(http.StatusNotFound, web.ResponseToClient(http.StatusNotFound, err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, web.ResponseToClient(http.StatusCreated, "Sukses membuat event", event))
}

func (controller *EventControllerImpl) GetEvent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("event_id"))

	getEvent, errGetEvent := controller.eventService.GetEvent(id, helper.JwtClaims{})
	if errGetEvent != nil {
		return c.JSON(http.StatusNotFound, web.ResponseToClient(http.StatusNotFound, errGetEvent.Error(), nil))
	}

	return c.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "succes", getEvent))

}

func (controller *EventControllerImpl) GetEventAdmin(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("event_id"))

	getEvent, errGetEvent := controller.eventService.GetEvent(id, helper.GetClaimsValue(c))
	if errGetEvent != nil {
		return c.JSON(http.StatusNotFound, web.ResponseToClient(http.StatusNotFound, errGetEvent.Error(), nil))
	}

	return c.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "succes", getEvent))

}

func (controller *EventControllerImpl) UpdateEvent(c echo.Context) error {
	event := new(web.EventUpdateServiceRequest)
	id, _ := strconv.Atoi(c.Param("event_id"))

	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	eventUpdate, errEventUpdate := controller.eventService.UpdateEvent(*event, id)

	if errEventUpdate != nil {
		return c.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, errEventUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "Event berhasil diupdate", eventUpdate))
}

func (controller *EventControllerImpl) DeleteEvent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("event_id"))

	err := controller.eventService.DeleteEvent(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "Event Berhasil Dihapus", nil))
}

func (controller *EventControllerImpl) GetAllEvents(e echo.Context) error {
	limit, _ := strconv.Atoi(e.QueryParam("limit"))
	page, _ := strconv.Atoi(e.QueryParam("page"))
	priceMax, _ := strconv.Atoi(e.QueryParam("price_max"))
	priceMin, _ := strconv.Atoi(e.QueryParam("price_min"))
	city := e.QueryParam("city")
	date := e.QueryParam("date")
	categoryId, _ := strconv.Atoi(e.QueryParam("category_id"))
	filter := e.QueryParam("filter")

	event := web.AllEventDataRequest{
		PriceMax:   priceMax,
		PriceMin:   priceMin,
		City:       city,
		Date:       date,
		CategoryId: categoryId,
		Filter:     web.Filter(filter),
		Limit:      limit,
		Page:       page,
	}

	response, err := controller.eventService.GetAllEvent(event)
	if err != nil {
		return e.JSON(http.StatusBadRequest, web.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return e.JSON(http.StatusOK, web.ResponseToClient(http.StatusOK, "sukses", response))
}
