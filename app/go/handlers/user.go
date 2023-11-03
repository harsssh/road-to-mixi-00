package handlers

import (
	"errors"
	"net/http"
	"problem1/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *services.UserService
}

type FriendInfo struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetFriendList(c echo.Context) error {
	uidString := c.QueryParam("id")
	uid, err := strconv.Atoi(uidString)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}

	friends, err := h.service.GetFriendList(uid)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return c.String(http.StatusNotFound, "user not found")
		} else {
			return c.String(http.StatusInternalServerError, "internal server error")
		}
	}

	friendInfoList := make([]*FriendInfo, len(friends))
	for i, friend := range friends {
		friendInfoList[i] = &FriendInfo{
			UserID: friend.UserID,
			Name:   friend.Name,
		}
	}
	return c.JSON(http.StatusOK, friendInfoList)
}

func (h *UserHandler) GetFriendOfFriendList(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *UserHandler) GetFriendOfFriendListPaging(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
