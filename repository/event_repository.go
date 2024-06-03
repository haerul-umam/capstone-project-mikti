package repository

import "github.com/haerul-umam/capstone-project-mikti/model/domain"

type EventRepository interface {
	GetEvent(Id int) (domain.Event, error)
	DecreaseQouta(event domain.Event) (domain.Event, error)
}
