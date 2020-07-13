package database

import (
	"database/sql"
	"fmt"

	"github.com/jdxj/study-web/config"

	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dsnFormat = "%s:%s@tcp(%s)/%s?parseTime=True&loc=Local"

	db *sql.DB
)

func init() {
	dbCfg := config.GetDB()
	dsn := fmt.Sprintf(dsnFormat, dbCfg.User, dbCfg.Pass, dbCfg.Host, dbCfg.Name)

	database, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err := database.Ping(); err != nil {
		panic(err)
	}
	db = database
}

func Close() {
	if err := db.Close(); err != nil {
		logs.Error("close db failed: %s", err)
	}
}
