package service

import (
	"github.com/haerul-umam/capstone-project-mikti/model/entity"
	"github.com/haerul-umam/capstone-project-mikti/model/web"
)

type CategoryService interface {
	CreateCategory(name string) (web.CategoryResponse, error)
	GetCategoryList() ([]entity.CategoryEntity, error)
	UpdateCategory(request web.CategoryUpdateServiceRequest, pathId int) (web.CategoryResponse, error)
}
