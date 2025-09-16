package repository

import (
	"context"
	"ecommerce/models"
	"errors"

	"gorm.io/gorm"
)

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := repo.Database.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func (r *UserRepository) CreateNewUser(ctx context.Context, user *models.User) (int64, error) {
	err := r.Database.WithContext(ctx).Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (repo *UserRepository) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	err := repo.Database.WithContext(ctx).Omit("password").Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, nil
		}

		return nil, err
	}
	return &user, nil
}
