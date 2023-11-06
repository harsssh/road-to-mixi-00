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

func excludeUsers(users []*models.User, excluded []*models.User) []*models.User {
	excludedMap := make(map[uint64]bool)
	for _, user := range excluded {
		excludedMap[user.ID] = true
	}

	result := make([]*models.User, 0, len(users))
	for _, user := range users {
		if !excludedMap[user.ID] {
			result = append(result, user)
		}
	}
	return result
}

func (u *UserService) GetFriendList(userID int) ([]*models.User, error) {
	panic("implement me")
}

func (u *UserService) GetFriendOfFriendList(userID int) ([]*models.User, error) {
	panic("implement me")
}

func (u *UserService) GetFriendOfFriendListPaging(userID int, page int, limit int) ([]*models.User, error) {
	panic("implement me")
}
