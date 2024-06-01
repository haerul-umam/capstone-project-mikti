package repository

import (
	"errors"

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
	var existingUserData domain.User

	errEmail := repo.db.First(&existingUserData, "email = ?", user.Email).Error

	if errEmail == nil {
		return domain.User{}, errors.New("user sudah terdaftar")
	}

	err := repo.db.Create(&user).Error

	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}
