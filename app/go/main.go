package main

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"problem1/configs"
	"problem1/handlers"
	"problem1/repository"
	"problem1/services"
	"strconv"
	"time"

	go_mysql "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const maxRetries = 5
const retryDelay = 5 * time.Second

func getDBConnection(c configs.DBConfig) (*gorm.DB, error) {
	mysqlConfig := go_mysql.NewConfig()
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = fmt.Sprintf("%s:%d", c.Host, c.Port)
	mysqlConfig.DBName = c.Name
	mysqlConfig.User = c.User
	mysqlConfig.ParseTime = true
	mysqlConfig.Collation = "utf8mb4_unicode_ci"

	var err error
	for i := 0; i < maxRetries; i++ {
		db, err := gorm.Open(mysql.Open(mysqlConfig.FormatDSN()), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			return db, nil
		}
		log.Printf("failed to connet to database")
		time.Sleep(retryDelay)
	}
	return nil, err
}

func main() {
	conf := configs.Get()

	db, err := getDBConnection(conf.DB)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "minimal_sns_app")
	})
	e.GET("/get_friend_list", userHandler.GetFriendList)
	e.GET("/get_friend_of_friend_list", userHandler.GetFriendOfFriendList)
	e.GET("/get_friend_of_friend_list_paging", userHandler.GetFriendOfFriendListPaging)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.Server.Port)))
}
