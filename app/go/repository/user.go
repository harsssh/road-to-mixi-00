package repository

import (
	"github.com/jmoiron/sqlx"
	"problem1/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) FindFriendsByUserID(userID int) ([]*models.User, error) {
	panic("implement me")
}

func (u *UserRepository) FindBlockUsersByUserID(userID int) ([]*models.User, error) {
	panic("implement me")
}

// FindFriendsOfFriendsByUserID does not exclude blocked users etc.
func (u *UserRepository) FindFriendsOfFriendsByUserID(userID int) ([]*models.User, error) {
	panic("implement me")
}

func (u *UserRepository) FindFriendsOfFriendsExcludingSomeUsersByUserIDWithPagination(
	userID int, excludedUserIDs []int, page int, limit int,
) ([]*models.User, error) {
	panic("implement me")
}
