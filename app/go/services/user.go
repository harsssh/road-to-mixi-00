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

func (u *UserService) GetFriendList(userID int) ([]*models.User, error) {
	return u.repo.FindFriendsByUserID(userID)
}

func (u *UserService) getUsersToExcludeForFriendsOfFriends(userID int) ([]*models.User, error) {
	friends, err := u.repo.FindFriendsByUserID(userID)
	if err != nil {
		return nil, err
	}
	blocks, err := u.repo.FindBlockUsersByUserID(userID)
	if err != nil {
		return nil, err
	}
	return append(friends, blocks...), nil
}

func (u *UserService) GetFriendOfFriendList(userID int) ([]*models.User, error) {
	friendsOfFriends, err := u.repo.FindFriendsOfFriendsByUserID(userID)
	if err != nil {
		return nil, err
	}
	usersToExclude, err := u.getUsersToExcludeForFriendsOfFriends(userID)
	if err != nil {
		return nil, err
	}
	result := excludeUsers(friendsOfFriends, usersToExclude)
	return result, nil
}

func (u *UserService) GetFriendOfFriendListPaging(userID int, page int, limit int) ([]*models.User, error) {
	usersToExclude, err := u.getUsersToExcludeForFriendsOfFriends(userID)
	if err != nil {
		return nil, err
	}
	userIDsToExclude := make([]int, len(usersToExclude))
	for i, user := range usersToExclude {
		userIDsToExclude[i] = user.UserID
	}
	result, err := u.repo.FindFriendsOfFriendsExcludingSomeUsersByUserIDWithPagination(
		userID, userIDsToExclude, page, limit,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
