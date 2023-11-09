package services

import "problem1/models"

//go:generate go run github.com/matryer/moq -pkg mocks -skip-ensure -out ../mocks/user_repository_mock.go . IUserRepository
type IUserRepository interface {
	FindFriendsByID(id int64) ([]*models.User, error)
	FindBlockUsersByID(id int64) ([]*models.User, error)
	FindFriendsOfFriendsByID(id int64) ([]*models.User, error)
	FindFriendsOfFriendsExcludingSomeUsersByIDWithPagination(
		id int64, excludeIDs []int64, page int, limit int,
	) ([]*models.User, error)
}

//go:generate go run github.com/matryer/moq -pkg mocks -skip-ensure -out ../mocks/user_service_mock.go . IUserService
type IUserService interface {
	GetFriendList(id int64) ([]*models.User, error)
	GetFriendOfFriendList(id int64) ([]*models.User, error)
	GetFriendOfFriendListPaging(id int64, page int, limit int) ([]*models.User, error)
}
