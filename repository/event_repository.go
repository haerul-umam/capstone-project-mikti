package repository

import "github.com/haerul-umam/capstone-project-mikti/model/domain"

type EventRepository interface {
	GetEvent(Id int) (domain.Event, error)
	DecreaseQouta(event domain.Event) (domain.Event, error)
	UpdateEvent(event domain.Event) (domain.Event, error)
	DeleteEvent(Id int) error
	GetAllEvent(priceMax int, priceMin int, city string, date string, categoryId int, filter string, limit int, page int) ([]domain.Event, int64, error)
}
