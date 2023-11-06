package services

import "problem1/models"

//go:generate moq -pkg mocks -out ../mocks/user_repository.go . IUserRepository
type IUserRepository interface {
	FindFriendsByUserID(userID int) ([]*models.User, error)
	FindFriendsOfFriendsByUserID(userID int) ([]*models.User, error)
	FindFriendsOfFriendsPagingByUserID(userID int, page int, limit int) ([]*models.User, error)
}

//go:generate moq -pkg mocks -out ../mocks/user_service.go . IUserService
type IUserService interface {
	GetFriendsByUserID(userID int) ([]*models.User, error)
	GetFriendsOfFriendsByUserID(userID int) ([]*models.User, error)
	GetFriendsOfFriendsPagingByUserID(userID int, page int, limit int) ([]*models.User, error)
}
