package service

import (
	"github.com/haerul-umam/capstone-project-mikti/model/entity"
	"github.com/haerul-umam/capstone-project-mikti/model/web"
)

type CategoryService interface {
	CreateCategory(name string) (web.CategoryCreateResponse, error)
	GetCategoryList() ([]entity.CategoryEntity, error)
	UpdateCategory(request web.CategoryUpdateServiceRequest, pathId int) (map[string]interface{}, error)
}
