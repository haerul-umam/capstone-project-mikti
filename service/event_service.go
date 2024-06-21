package service

import (
	"github.com/haerul-umam/capstone-project-mikti/helper"
	"github.com/haerul-umam/capstone-project-mikti/model/web"
)

type EventService interface {
	CreateEvent(request web.EventCreateServiceRequest) (web.EventUpdateCreateResponse, error)
	UpdateEvent(request web.EventUpdateServiceRequest, pathID int) (interface{}, error)
	GetEvent(eventId int, user helper.JwtClaims) (interface{}, error)
	DeleteEvent(pathId int) error
	GetAllEvent(request web.AllEventDataRequest) (web.AllEventDataResponse, error)
}
