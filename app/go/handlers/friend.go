package handlers

import (
	"net/http"
	"problem1/services"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
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
