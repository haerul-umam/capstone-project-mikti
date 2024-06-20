package service

import (
	"errors"
	"math"

	"github.com/haerul-umam/capstone-project-mikti/helper"
	"github.com/haerul-umam/capstone-project-mikti/model/domain"
	"github.com/haerul-umam/capstone-project-mikti/model/entity"
	"github.com/haerul-umam/capstone-project-mikti/model/web"
	"github.com/haerul-umam/capstone-project-mikti/repository"
)

type EventServiceImpl struct {
	repository repository.EventRepository
}

func NewEventService(repository repository.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{
		repository: repository,
	}
}

func (service *EventServiceImpl) GetEvent(eventId int, user helper.JwtClaims) (interface{}, error) {

	getEvent, errGetEvent := service.repository.GetEvent(eventId)

	if user.ID != "" {

		if errGetEvent != nil {
			return web.EventDetailResponseAdmin{}, errGetEvent
		}

		return web.EventDetailResponseAdmin{
			ItemID:      getEvent.EventID,
			CategoryID:  getEvent.CategoryID,
			Name:        getEvent.Name,
			Date:        getEvent.Date,
			Price:       getEvent.Price,
			Is_free:     getEvent.Is_free,
			City:        getEvent.City,
			Description: getEvent.Description,
			Quota:       getEvent.Quota,
			DeletedAt:   getEvent.DeletedAt,
			Category: web.Category{
				Id:   getEvent.Category.ID,
				Name: getEvent.Category.Name,
			},
		}, nil

	} else if getEvent.DeletedAt.Valid {
		return web.EventDetailResponse{}, errors.New("event tidak ditemukan")

	} else {
		return web.EventDetailResponse{
			ItemID:      getEvent.EventID,
			CategoryID:  getEvent.CategoryID,
			Name:        getEvent.Name,
			Date:        getEvent.Date,
			Price:       getEvent.Price,
			Is_free:     getEvent.Is_free,
			City:        getEvent.City,
			Description: getEvent.Description,
			Quota:       getEvent.Quota,
			Category: web.Category{
				Id:   getEvent.Category.ID,
				Name: getEvent.Category.Name,
			},
		}, nil

	}
}

func (service *EventServiceImpl) UpdateEvent(request web.EventUpdateServiceRequest, pathID int) (interface{}, error) {
	getEventById, err := service.repository.GetEvent(pathID)

	if err != nil {
		return getEventById, err
	}

	eventRequest := domain.Event{
		EventID:     getEventById.EventID,
		Name:        request.Name,
		CategoryID:  request.CategoryID,
		Date:        request.Date,
		Price:       request.Price,
		Is_free:     request.Is_free,
		City:        request.City,
		Description: request.Description,
		Quota:       request.Quota,
	}

	eventUpdate, errUpdate := service.repository.UpdateEvent(eventRequest)

	if errUpdate != nil {
		return entity.ToEventEntity(eventUpdate), errUpdate
	}

	return web.EventUpdateResponse{
		ItemID:      getEventById.EventID,
		CategoryID:  request.CategoryID,
		Name:        request.Name,
		Date:        request.Date,
		Price:       request.Price,
		Is_free:     request.Is_free,
		City:        request.City,
		Description: request.Description,
		Quota:       request.Quota,
	}, errUpdate
}

func (service *EventServiceImpl) DeleteEvent(pathId int) error {
	err := service.repository.DeleteEvent(pathId)

	if err != nil {
		return err
	}

	return nil
}

func (service *EventServiceImpl) GetAllEvent(request web.AllEventDataRequest) (web.AllEventDataResponse, error) {
	eventReq := web.AllEventDataRequest{
		PriceMax:   request.PriceMax,
		PriceMin:   request.PriceMin,
		City:       request.City,
		Date:       request.Date,
		CategoryId: request.CategoryId,
		Filter:     request.Filter,
		Limit:      request.Limit,
		Page:       request.Page,
	}

	if eventReq.Limit <= 0 {
		eventReq.Limit = 10
	}

	if eventReq.Page <= 0 {
		eventReq.Page = 1
	}

	getEvents, total, errGetEvents := service.repository.GetAllEvent(eventReq.PriceMax, eventReq.PriceMin, eventReq.City, eventReq.Date, eventReq.CategoryId, string(eventReq.Filter), eventReq.Limit, eventReq.Page)

	if errGetEvents != nil {
		return web.AllEventDataResponse{}, errGetEvents
	}

	totalPages := int(math.Ceil(float64(total) / float64(eventReq.Limit)))
	eventEntities := entity.ToEventEntities(getEvents)

	return web.AllEventDataResponse{
		Total:       total,
		TotalPages:  totalPages,
		CurrentPage: eventReq.Page,
		Events:      eventEntities,
	}, nil
}
