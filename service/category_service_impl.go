package service

import (
	"github.com/haerul-umam/capstone-project-mikti/helper"
	"github.com/haerul-umam/capstone-project-mikti/model/domain"
	"github.com/haerul-umam/capstone-project-mikti/model/entity"
	"github.com/haerul-umam/capstone-project-mikti/model/web"
	"github.com/haerul-umam/capstone-project-mikti/repository"
)

type CategoryServiceImpl struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{repo}
}

func (service *CategoryServiceImpl) CreateCategory(name string) (web.CategoryCreateResponse, error) {

	CategoryReq := domain.Category{
		Name: name,
	}

	saveCategory, errSaveCategory := service.repository.CreateCategory(CategoryReq)

	if errSaveCategory != nil {
		return web.CategoryCreateResponse{}, errSaveCategory
	}

	return web.CategoryCreateResponse{ID: saveCategory.ID, Name: saveCategory.Name}, nil
}

func (service *CategoryServiceImpl) GetCategoryList() ([]entity.CategoryEntity, error) {
	getCategoryList, errGetCategoryList := service.repository.GetCategories()

	if errGetCategoryList != nil {
		return []entity.CategoryEntity{}, errGetCategoryList
	}

	return entity.ToCategoryEntities(getCategoryList), nil
}

func (service *CategoryServiceImpl) UpdateCategory(request web.CategoryUpdateServiceRequest, pathId int) (map[string]interface{}, error) {
	categoryRequest := domain.Category{
		ID:   pathId,
		Name: request.Name,
	}

	UpdateCategory, errCategory := service.repository.UpdateCategory(categoryRequest)

	if errCategory != nil {
		return nil, errCategory
	}

	return helper.ResponseToJson{"name": UpdateCategory.Name}, nil
}
