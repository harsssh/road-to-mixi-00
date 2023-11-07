package services

import "problem1/models"

//go:generate go run github.com/matryer/moq -pkg mocks -skip-ensure -out ../mocks/user_repository_mock.go . IUserRepository
type IUserRepository interface {
	FindFriendsByUserID(userID int) ([]*models.User, error)
	FindBlockUsersByUserID(userID int) ([]*models.User, error)
	FindFriendsOfFriendsByUserID(userID int) ([]*models.User, error)
	FindFriendsOfFriendsExcludingSomeUsersByUserIDWithPagination(
		userID int, excludedUserIDs []int, page int, limit int,
	) ([]*models.User, error)
}

//go:generate go run github.com/matryer/moq -pkg mocks -skip-ensure -out ../mocks/user_service_mock.go . IUserService
type IUserService interface {
	GetFriendList(userID int) ([]*models.User, error)
	GetFriendOfFriendList(userID int) ([]*models.User, error)
	GetFriendOfFriendListPaging(userID int, page int, limit int) ([]*models.User, error)
}
