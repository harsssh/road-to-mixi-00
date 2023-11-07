package repository

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"problem1/configs"
	"time"
)

const maxRetries = 5
const retryDelay = 5 * time.Second

func ConnectToDB(c configs.DBConfig) (*sqlx.DB, error) {
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = fmt.Sprintf("%s:%d", c.Host, c.Port)
	mysqlConfig.DBName = c.Name
	mysqlConfig.User = c.User
	mysqlConfig.ParseTime = true

	for i := 0; i < maxRetries; i++ {
		db, err := sqlx.Connect("mysql", mysqlConfig.FormatDSN())
		if err == nil {
			return db, nil
		}
		time.Sleep(retryDelay)
	}
	return nil, fmt.Errorf("failed to connect to db: %s", mysqlConfig.FormatDSN())
}
