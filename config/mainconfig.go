package config

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nsplnpbjy/bbs/midware"
	"github.com/nsplnpbjy/bbs/service"
)

func InitEngine(path string) *gin.Engine {
	r := gin.Default()
	//在这里插入中间件
	r.Use(midware.RateLimitMiddleWare(time.Second, 100, 100))
	r.Use(midware.LogMidware())
	cs := NewControllerSet()
	//在这里插入controller
	cs.ADD("/hello", Method_GET, service.HelloWorld)
	SetRouter(r, path, cs)
	return r
}
