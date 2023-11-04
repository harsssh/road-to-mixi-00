package repository

import (
	"gorm.io/gorm"
	"problem1/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) FindFriendsByUserID(userID int) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) FindFriendsOfFriendsByUserID(userID int) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) FindFriendsOfFriendsPagingByUserID(userID int, page int, limit int) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}
