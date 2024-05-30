package repository

import (
	"github.com/haerul-umam/capstone-project-mikti/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (repo *UserRepositoryImpl) GetUserByEmail(email string) (*domain.User, error) {
	user := new(domain.User)

	if err := repo.db.Where("email = ?", email).Take(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) SaveUser(user domain.User) (domain.User, error) {
	err := repo.db.Create(&user).Error

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
