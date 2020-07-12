package routers

import (
	"net/http"

	"github.com/jdxj/study-web/config"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
)

func NewServer() *http.Server {
	handler := newRouter()

	serverCfg := config.GetServer()
	server := &http.Server{
		Addr:    serverCfg.Addr,
		Handler: handler,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logs.Error("listen and server err: %s", err)
		}
	}()
	return server
}

func newRouter() *gin.Engine {
	engine := gin.New()

	engine.Use(gin.Recovery())
	engine.Use(RecordUserAgent())

	engine.GET("/", SayHello)

	return engine
}

func SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
