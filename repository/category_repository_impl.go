package repository

import (
	"errors"

	"github.com/haerul-umam/capstone-project-mikti/model/domain"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{db: db}
}

func (repo *CategoryRepositoryImpl) CreateCategory(category domain.Category) (domain.Category, error) {
	err := repo.db.Create(&category).Error

	if err != nil {
		return domain.Category{}, nil
	}
	return category, nil
}

func (repo *CategoryRepositoryImpl) GetCategories() ([]domain.Category, error) {
	var categories []domain.Category

	err := repo.db.Find(&categories).Error

	if err != nil {
		return []domain.Category{}, err
	}

	return categories, nil
}

func (repo *CategoryRepositoryImpl) UpdateCategory(category domain.Category) (domain.Category, error) {
	var categoryData domain.Category
	err := repo.db.First(&categoryData, "id = ?", category.ID).Error
	if err != nil {
		return domain.Category{}, errors.New("category tidak di temukan")
	}

	err = repo.db.Model(&categoryData).Where("id = ?", category.ID).Updates(category).Error

	if err != nil {
		return category, err
	}
	return category, nil
}

func (repo *CategoryRepositoryImpl) DeleteCategory(categoryID uint) error {
	var category domain.Category
	err := repo.db.First(&category, "id = ?", categoryID).Error
	if err != nil {
		return errors.New("category tidak ditemukan")
	}

	err = repo.db.Delete(&category).Error
	if err != nil {
		return err
	}

	return nil
}
