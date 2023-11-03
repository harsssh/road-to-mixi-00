package tests

import (
	"problem1/models"
	"problem1/services"
)

type MockUserService struct {
	data []*models.User
}

type MockUserRepository struct {
	data map[int]*models.User
}

func NewMockUserService(data []*models.User) *MockUserService {
	return &MockUserService{data: data}
}

func NewMockUserRepository(data map[int]*models.User) *MockUserRepository {
	return &MockUserRepository{data: data}
}

func (s *MockUserService) GetFriendList(_ int) ([]*models.User, error) {
	return s.data, nil
}

func (r *MockUserRepository) FindByUserID(uid int) (*models.User, error) {
	if user, ok := r.data[uid]; ok {
		return user, nil
	}
	return nil, services.ErrUserNotFound
}
