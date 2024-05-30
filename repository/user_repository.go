package repository

import "github.com/haerul-umam/capstone-project-mikti/model/domain"

type UserRepository interface {
	GetUserByEmail(email string) (*domain.User, error)
	SaveUser(user domain.User) (domain.User, error)
}
