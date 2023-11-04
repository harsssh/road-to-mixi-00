package tests

import (
	"problem1/models"
	"problem1/services"
)

type MockUserService struct {
	data map[int]*models.User
}

type MockUserRepository struct {
	data map[int]*models.User
}

func NewMockUserService(data map[int]*models.User) *MockUserService {
	return &MockUserService{data: data}
}

func NewMockUserRepository(data map[int]*models.User) *MockUserRepository {
	return &MockUserRepository{data: data}
}

func (s *MockUserService) GetFriendList(uid int) ([]*models.User, error) {
	if user, ok := s.data[uid]; ok {
		return user.Friends, nil
	}
	return nil, services.ErrUserNotFound
}

func (s *MockUserService) GetFriendOfFriendList(uid int) ([]*models.User, error) {
	var user *models.User
	for _, u := range s.data {
		if u.UserID == uid {
			user = u
			break
		}
	}
	if user == nil {
		return nil, services.ErrUserNotFound
	}

	var friends []*models.User
	for _, friend := range user.Friends {
		friends = append(friends, friend.Friends...)
	}
	return friends, nil
}

func (r *MockUserRepository) FindByUserIDWithFriends(uid int) (*models.User, error) {
	if user, ok := r.data[uid]; ok {
		return user, nil
	}
	return nil, services.ErrUserNotFound
}

func (r *MockUserRepository) FindFriendsByUserIDWithFriends(uid int) ([]*models.User, error) {
	if user, ok := r.data[uid]; ok {
		return user.Friends, nil
	}
	return nil, services.ErrUserNotFound
}
