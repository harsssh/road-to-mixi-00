package services

import (
	"problem1/models"
)

type UserService struct {
	repo IUserRepository
}

func NewUserService(r IUserRepository) *UserService {
	return &UserService{repo: r}
}

func (u *UserService) GetFriendList(userID int) ([]*models.User, error) {
	friends, err := u.repo.FindFriendsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (u *UserService) GetFriendOfFriendList(userID int) ([]*models.User, error) {
	friends, err := u.repo.FindFriendsOfFriendsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (u *UserService) GetFriendOfFriendListPaging(userID int, page int, limit int) ([]*models.User, error) {
	friends, err := u.repo.FindFriendsOfFriendsPagingByUserID(userID, page, limit)
	if err != nil {
		return nil, err
	}
	return friends, nil
}
