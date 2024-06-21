package repository

import (
	"errors"

	"github.com/haerul-umam/capstone-project-mikti/model/domain"
	"github.com/haerul-umam/capstone-project-mikti/model/web"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{db}
}

func (repo *EventRepositoryImpl) CreateEvent(event domain.Event) (domain.Event, error) {
	err := repo.db.Create(&event).Error

	if err != nil {
		return domain.Event{}, err
	}

	return event, nil
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

func (repo *EventRepositoryImpl) GetAllEvent(priceMax int, priceMin int, city string, date string, categoryId int, filter string, limit int, page int) ([]domain.Event, int64, error) {
	var events []domain.Event
	allEventData := repo.db.Model(&domain.Event{}).Preload("Category")

	if priceMax > 0 {
		errPriceMax := allEventData.Where("event.price <= ?", priceMax).Error
		if errPriceMax != nil {
			return []domain.Event{}, 0, errPriceMax
		}
	}

	if priceMin > 0 {
		errPriceMin := allEventData.Where("event.price >= ?", priceMin).Error
		if errPriceMin != nil {
			return []domain.Event{}, 0, errPriceMin
		}
	}

	if city != "" {
		errCity := allEventData.Where("event.city = ?", city).Error
		if errCity != nil {
			return []domain.Event{}, 0, errCity
		}
	}

	if date != "" {
		errDate := allEventData.Where("event.date = ?", date).Error
		if errDate != nil {
			return []domain.Event{}, 0, errDate
		}
	}

	if categoryId > 0 {
		errCategoryId := allEventData.Where("event.category_id = ?", categoryId).Error
		if errCategoryId != nil {
			return []domain.Event{}, 0, errCategoryId
		}
	}

	switch filter {
	case string(web.Termurah):
		errFilterTermurah := allEventData.Order("event.price ASC").Error
		if errFilterTermurah != nil {
			return []domain.Event{}, 0, errFilterTermurah
		}
	case string(web.Termahal):
		errFilterTermahal := allEventData.Order("event.price DESC").Error
		if errFilterTermahal != nil {
			return []domain.Event{}, 0, errFilterTermahal
		}
	case string(web.Terbaru):
		errFilterTerbaru := allEventData.Order("event.date DESC").Error
		if errFilterTerbaru != nil {
			return []domain.Event{}, 0, errFilterTerbaru
		}
	case string(web.Terpopuler):
		errFilterTerpopuler := allEventData.
			Select(`event.id, event.category_id, event.name, event.date, event.price, event.is_free, event.city, event.description, event.quota, COUNT("order".event_id) AS jumlah_pesanan`).
			Joins(`LEFT JOIN "order" ON "order".event_id = event.id`).
			Group("event.id").
			Order(`COUNT("order".event_id) DESC, event.id ASC`).Error
		if errFilterTerpopuler != nil {
			return []domain.Event{}, 0, errFilterTerpopuler
		}

	}

	var total int64
	err := allEventData.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	err = allEventData.Limit(limit).Offset(offset).Find(&events).Error
	if err != nil {
		return nil, 0, err
	}

	return events, total, nil

}
