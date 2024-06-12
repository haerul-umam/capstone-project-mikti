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

	err := repo.db.Unscoped().Model(&domain.Event{}).Preload("Category").First(&eventData, "id= ?", Id).Error

	if err != nil {
		return domain.Event{}, errors.New("event tidak ditemukan")
	}

	return eventData, nil
}

func (repo *EventRepositoryImpl) DecreaseQouta(event domain.Event) (domain.Event, error) {
	err := repo.db.Model(domain.Event{}).Where("id = ?", event.EventID).Updates(event).Error

	if err != nil {
		return event, err
	}

	return event, nil
}
func (repo *EventRepositoryImpl) UpdateEvent(event domain.Event) (domain.Event, error) {
	err := repo.db.Model(domain.Event{}).Where("id = ?", event.EventID).Updates(event).Error

	if err != nil {
		return event, err
	}

	return event, nil
}

func (repo *EventRepositoryImpl) DeleteEvent(Id int) error {
	var eventData domain.Event

	err := repo.db.First(&eventData, "id = ?", Id).Error
	if err != nil {
		return err
	}

	err = repo.db.Delete(&eventData).Error
	if err != nil {
		return err
	}

	return nil
}
