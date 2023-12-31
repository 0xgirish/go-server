package main

import (
	"fmt"

	"github.com/0xgirish/go-server/configs"
	"github.com/0xgirish/go-server/repository"
	"github.com/0xgirish/go-server/server"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	config := configs.Get()

	db := getDbConn(config)
	defer db.Close()

	todos := repository.NewTodoRepository(db)

	s := server.New(todos)

	host := config.GetString("server.host")
	port := config.GetInt("server.port")

	s.Start(host, port)
}

func getDbConn(config *viper.Viper) *sqlx.DB {
	cfg := mysql.NewConfig()
	cfg.User = config.GetString("db.user")
	cfg.Passwd = config.GetString("db.password")
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%d", config.GetString("db.host"), config.GetInt("db.port"))
	cfg.DBName = config.GetString("db.name")

	dsn := cfg.FormatDSN()
	return sqlx.MustOpen("mysql", dsn)
}
