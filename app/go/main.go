package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"problem1/configs"
	"problem1/handlers"
	"problem1/repository"
	"problem1/services"
	"strconv"
)

func main() {
	conf := configs.Get()

	db, err := repository.ConnectToDB(conf.DB)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "minimal_sns_app")
	})
	e.GET("/get_friend_list", userHandler.GetFriendList)
	e.GET("/get_friend_of_friend_list", userHandler.GetFriendOfFriendList)
	e.GET("/get_friend_of_friend_list_paging", userHandler.GetFriendOfFriendListPaging)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.Server.Port)))
}
