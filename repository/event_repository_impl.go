package repository

import (
	"github.com/haerul-umam/capstone-project-mikti/model/domain"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{db: db}
}

func (repo *EventRepositoryImpl) GetEvent(Id int) (domain.Event, error) {
	var eventData domain.Event
	err := repo.db.Table("event").First(&eventData, "id = ?", Id).Error

	if err != nil {
		return domain.Event{}, err
	}

	return eventData, nil
}

func (repo *EventRepositoryImpl) UpdateEvent(event domain.Event) (domain.Event, error) {
	err := repo.db.Model(domain.Event{}).Where("id = ?", event.EventID).Updates(event).Error

	if err != nil {
		return event, err
	}

	return event, nil
}
