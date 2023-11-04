package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"problem1/services"
)

type UserHandler struct {
	service services.IUserService
}

type FriendInfo struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

func NewUserHandler(s services.IUserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetFriendList(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *UserHandler) GetFriendOfFriendList(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *UserHandler) GetFriendOfFriendListPaging(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
