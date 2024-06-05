package service

import (
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

func (service *EventServiceImpl) GetEvent(eventId int) (entity.EventEntity, error) {
	getEvent, errGetEvent := service.repository.GetEvent(eventId)

	if errGetEvent != nil {
		return entity.EventEntity{}, errGetEvent
	}

	return entity.ToEventEntity(getEvent), nil
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

	return entity.ToEventEntity(eventUpdate), errUpdate
}
