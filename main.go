package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/jdxj/study-web/routers"
)

func init() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"study-web.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func main() {
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
	logs.Info("shutdown")
	logs.GetBeeLogger().Close()
}
