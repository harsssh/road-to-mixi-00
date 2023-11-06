package handlers

import (
	"github.com/labstack/echo/v4"
	"problem1/models"
	"strconv"
)

type FriendListEntry struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

func getUserID(c echo.Context) (int, error) {
	userID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func getPaginationParams(c echo.Context) (PaginationParams, error) {
	var page int
	var err error
	if c.QueryParam("page") == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			return PaginationParams{}, err
		}
	}

	var limit int
	if c.QueryParam("limit") == "" {
		limit = defaultLimit
	} else {
		limit, err = strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			return PaginationParams{}, err
		}
	}

	return PaginationParams{
		Page:  page,
		Limit: limit,
	}, nil
}

func convertToFriendList(users []*models.User) []*FriendListEntry {
	friends := make([]*FriendListEntry, len(users))
	for i, user := range users {
		friends[i] = &FriendListEntry{
			UserID: user.UserID,
			Name:   user.Name,
		}
	}
	return friends
}
