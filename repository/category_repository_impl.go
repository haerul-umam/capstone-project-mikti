package repository

import (
	"errors"
	"time"

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

	err := repo.db.Where("deleted_at IS NULL").Find(&categories).Error

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

func (repo *CategoryRepositoryImpl) DeleteCategory(id int) error {
	var category domain.Category
	result := repo.db.First(&category, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errors.New("category Sudah Di Hapus")
		}
		return result.Error
	}
	now := time.Now()
	return repo.db.Model(&domain.Category{}).Where("id = ?", id).Update("deleted_at", &now).Error
}
