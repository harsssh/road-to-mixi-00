package services

import "problem1/models"

type IUserRepository interface {
	FindByUserIDWithFriends(int) (*models.User, error)
	FindFriendsByUserIDWithFriends(int) ([]*models.User, error)
}

type IUserService interface {
	GetFriendList(int) ([]*models.User, error)
	GetFriendOfFriendList(int) ([]*models.User, error)
}

type UserService struct {
	repo IUserRepository
}

func NewUserService(r IUserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetFriendList(uid int) ([]*models.User, error) {
	user, err := s.repo.FindByUserIDWithFriends(uid)
	if err != nil {
		return nil, err
	}
	return user.Friends, nil
}

func (s *UserService) GetFriendOfFriendList(uid int) ([]*models.User, error) {
	user, err := s.repo.FindFriendsByUserIDWithFriends(uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
