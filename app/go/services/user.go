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
	//TODO implement me
	panic("implement me")
}

func (u *UserService) GetFriendOfFriendList(userID int) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) GetFriendOfFriendListPaging(userID int, page int, limit int) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}
