package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jdxj/study-web/config"

	"github.com/jdxj/study-web/database"
	"github.com/jdxj/study-web/routers"

	"github.com/astaxie/beego/logs"
)

const logConfig = `{
    "filename":"study-web.log",
    "level":7,
    "maxlines":0,
    "maxsize":0,
    "daily":true,
    "maxdays":10,
    "color":true
}`

func init() {
	loggerCfg := config.GetLogger()
	err := logs.SetLogger(logs.AdapterFile, loggerCfg.String())
	if err != nil {
		panic(err)
	}
}

func main() {
	logs.Info("start")
	server := routers.NewServer()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logs.Info("stopping")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logs.Error("shutdown server err: %s", err)
	}
	<-ctx.Done()
	database.CloseDB()
	logs.Info("shutdown")
	logs.GetBeeLogger().Close()
}
