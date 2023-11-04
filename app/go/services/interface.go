package services

import "problem1/models"

type IUserRepository interface {
	FindFriendsByUserID(userID int) ([]*models.User, error)
	FindFriendsOfFriendsByUserID(userID int) ([]*models.User, error)
	FindFriendsOfFriendsPagingByUserID(userID int, page int, limit int) ([]*models.User, error)
}

type IUserService interface {
	GetFriendsByUserID(userID int) ([]*models.User, error)
	GetFriendsOfFriendsByUserID(userID int) ([]*models.User, error)
	GetFriendsOfFriendsPagingByUserID(userID int, page int, limit int) ([]*models.User, error)
}

//go:generate bulkmockgen -use_go_run ISet -- -package mocks -destination ../mocks/mock.go
var ISet = []any{
	new(IUserRepository),
	new(IUserService),
}
