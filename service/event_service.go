package service

import (
	"github.com/haerul-umam/capstone-project-mikti/model/web"
)

type EventService interface {
	UpdateEvent(request web.EventUpdateServiceRequest, pathID int) (interface{}, error)
	GetEvent(eventId int) (interface{}, error)
}
