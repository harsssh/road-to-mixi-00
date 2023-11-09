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
	excludedMap := make(map[int64]bool)
	for _, user := range excluded {
		excludedMap[user.ID] = true
	}

	result := make([]*models.User, 0, len(users))
	for _, user := range users {
		if _, ok := excludedMap[user.ID]; !ok {
			result = append(result, user)
		}
	}
	return result
}

func (u *UserService) GetFriendList(id int64) ([]*models.User, error) {
	return u.repo.FindFriendsByID(id)
}

func (u *UserService) getUsersToExcludeForFriendsOfFriends(id int64) ([]*models.User, error) {
	friends, err := u.repo.FindFriendsByID(id)
	if err != nil {
		return nil, err
	}
	blocks, err := u.repo.FindBlockUsersByID(id)
	if err != nil {
		return nil, err
	}
	return append(friends, blocks...), nil
}

func (u *UserService) GetFriendOfFriendList(id int64) ([]*models.User, error) {
	friendsOfFriends, err := u.repo.FindFriendsOfFriendsByID(id)
	if err != nil {
		return nil, err
	}
	usersToExclude, err := u.getUsersToExcludeForFriendsOfFriends(id)
	if err != nil {
		return nil, err
	}
	result := excludeUsers(friendsOfFriends, usersToExclude)
	return result, nil
}

func (u *UserService) GetFriendOfFriendListPaging(id int64, page int, limit int) ([]*models.User, error) {
	usersToExclude, err := u.getUsersToExcludeForFriendsOfFriends(id)
	if err != nil {
		return nil, err
	}
	excludeIDs := make([]int64, len(usersToExclude))
	for i, user := range usersToExclude {
		excludeIDs[i] = user.ID
	}
	result, err := u.repo.FindFriendsOfFriendsExcludingSomeUsersByIDWithPagination(
		id, excludeIDs, page, limit,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
