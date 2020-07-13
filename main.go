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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logs.Error("shutdown server err: %s", err)
	}
	<-ctx.Done()
	database.Close()

	logs.Info("shutdown")
	logs.GetBeeLogger().Close()
}
