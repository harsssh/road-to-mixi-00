package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"problem1/services"
)

type UserHandler struct {
	service services.IUserService
}

func NewUserHandler(s services.IUserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetFriendList(c echo.Context) error {
	userID, err := getUserID(c)
	if err != nil || userID < 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	friends, err := h.service.GetFriendList(userID)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return c.NoContent(http.StatusNotFound)
		}
		return c.NoContent(http.StatusInternalServerError)
	}
	friendList := convertToFriendList(friends)
	return c.JSON(http.StatusOK, friendList)
}

func (h *UserHandler) GetFriendOfFriendList(c echo.Context) error {
	userID, err := getUserID(c)
	if err != nil || userID < 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	friends, err := h.service.GetFriendOfFriendList(userID)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return c.NoContent(http.StatusNotFound)
		}
		return c.NoContent(http.StatusInternalServerError)
	}
	friendList := convertToFriendList(friends)
	return c.JSON(http.StatusOK, friendList)
}

func (h *UserHandler) GetFriendOfFriendListPaging(c echo.Context) error {
	userID, err := getUserID(c)
	if err != nil || userID < 0 {
		return c.NoContent(http.StatusBadRequest)
	}
	params, err := getPaginationParams(c)
	if err != nil || !params.isValid() {
		return c.NoContent(http.StatusBadRequest)
	}

	friends, err := h.service.GetFriendOfFriendListPaging(userID, params.Page, params.Limit)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return c.NoContent(http.StatusNotFound)
		}
		return c.NoContent(http.StatusInternalServerError)
	}
	friendList := convertToFriendList(friends)
	return c.JSON(http.StatusOK, friendList)
}
