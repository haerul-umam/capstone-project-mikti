package service

import (
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

func (service *CategoryServiceImpl) CreateCategory(name string) (web.CategoryResponse, error) {

	CategoryReq := domain.Category{
		Name: name,
	}

	saveCategory, errSaveCategory := service.repository.CreateCategory(CategoryReq)

	if errSaveCategory != nil {
		return web.CategoryResponse{}, errSaveCategory
	}

	return web.CategoryResponse{ID: saveCategory.ID, Name: saveCategory.Name}, nil
}

func (service *CategoryServiceImpl) GetCategoryList() ([]entity.CategoryEntity, error) {
	getCategoryList, errGetCategoryList := service.repository.GetCategories()

	if errGetCategoryList != nil {
		return []entity.CategoryEntity{}, errGetCategoryList
	}

	return entity.ToCategoryEntities(getCategoryList), nil
}

func (service *CategoryServiceImpl) UpdateCategory(request web.CategoryUpdateServiceRequest, pathId int) (web.CategoryResponse, error) {
	categoryRequest := domain.Category{
		ID:   pathId,
		Name: request.Name,
	}

	updateCategory, errCategory := service.repository.UpdateCategory(categoryRequest)

	if errCategory != nil {
		return web.CategoryResponse{}, errCategory
	}

	return web.CategoryResponse{ID: updateCategory.ID, Name: updateCategory.Name}, nil
}

func (service *CategoryServiceImpl) DeleteCategory(pathId uint) error {
	err := service.repository.DeleteCategory(uint(pathId))

	if err != nil {
		return err
	}

	return nil
}
