package entity

import (
	"github.com/haerul-umam/capstone-project-mikti/model/domain"
)

type UserEntity struct {
	UserID string      `json:"id"`
	Name   string      `json:"name"`
	Email  string      `json:"email"`
	Role   domain.Role `json:"role"`
}

func ToUserEntity(id, name, email string, role domain.Role) UserEntity {
	return UserEntity{id, name, email, role}
}

func ToUserEntities(users []domain.User) []UserEntity {
	data := []UserEntity{}

	for _, user := range users {
		data = append(data, ToUserEntity(user.UserID, user.Name, user.Email, user.Role))
	}

	return data
}
