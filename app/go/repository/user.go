package repository

import (
	"errors"
	"problem1/models"
	"problem1/services"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindByUserIDWithFriends(uid int) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Friends").First(&user, "user_id = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, services.ErrUserNotFound
		}
	}
	return &user, nil
}

func (r *UserRepository) FindFriendsByUserIDWithFriends(uid int) ([]*models.User, error) {
	var user models.User
	if err := r.db.Preload("Friends.Friends").First(&user, "user_id = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, services.ErrUserNotFound
		}
	}
	return user.Friends, nil
}
