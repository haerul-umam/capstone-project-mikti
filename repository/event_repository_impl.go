package repository

import (
	"errors"

	"github.com/haerul-umam/capstone-project-mikti/model/domain"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{db}
}

func (repo *EventRepositoryImpl) GetEvent(Id int) (domain.Event, error) {
	var eventData domain.Event

	err := repo.db.First(&eventData, "id = ?", Id).Error

	if err != nil {
		return domain.Event{}, errors.New("event tidak ditemukan")
	}

	return eventData, nil
}
