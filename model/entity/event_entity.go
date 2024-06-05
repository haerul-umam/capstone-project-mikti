package entity

import "github.com/haerul-umam/capstone-project-mikti/model/domain"

type EventEntity struct {
	EventID     int    `json:"event_id"`
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	Price       int    `json:"price"`
	Is_free     bool   `json:"is_free"`
	City        string `json:"city"`
	Description string `json:"description"`
	Quota       int    `json:"quota"`
}

func ToEventEntity(event domain.Event) EventEntity {

	return EventEntity{
		EventID:     event.EventID,
		CategoryID:  event.CategoryID,
		Name:        event.Name,
		Date:        event.Date,
		Price:       event.Price,
		Is_free:     event.Is_free,
		City:        event.City,
		Description: event.Description,
		Quota:       event.Quota,
	}
}

// func ToEventEntities(event []domain.Event) []EventEntity {
// 	data := []EventEntity{}

// 	for _, event := range events {
// 		data = appe
// 	}

// 	return data
// }
