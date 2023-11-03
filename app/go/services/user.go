package services

import "problem1/models"

type IUserRepository interface {
	FindByUserID(int) (*models.User, error)
}

type IUserService interface {
	GetFriendList(int) ([]*models.User, error)
}

type UserService struct {
	repo IUserRepository
}

func NewUserService(r IUserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetFriendList(uid int) ([]*models.User, error) {
	user, err := s.repo.FindByUserID(uid)
	if err != nil {
		return nil, err
	}
	return user.Friends, nil
}
