package repository

import "github.com/haerul-umam/capstone-project-mikti/model/domain"

type CategoryRepository interface {
	CreateCategory(category domain.Category) (domain.Category, error)
	GetCategories() ([]domain.Category, error)
	UpdateCategory(category domain.Category) (domain.Category, error)
}
