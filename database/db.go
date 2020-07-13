package database

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/jdxj/study-web/config"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dsnFormat = "%s:%s@tcp(%s)/%s?parseTime=True&loc=Local"

	db *sql.DB
)

func init() {
	dbCfg := config.GetDB()
	dsn := fmt.Sprintf(dsnFormat, dbCfg.User, dbCfg.Pass, dbCfg.Host, dbCfg.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	db = db
}

func CloseDB() {
	if err := db.Close(); err != nil {
		logs.Error("close db failed: %s", err)
	}
}
