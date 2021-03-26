package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"goBlog/config"
)

var db *sqlx.DB

func init() {
	db = sqlx.MustConnect(config.Mysql, config.GetString("database.enable"))
}
