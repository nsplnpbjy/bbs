package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nsplnpbjy/bbs/config"
	"github.com/nsplnpbjy/bbs/midware"
)

func InitEngine(path string) *gin.Engine {
	r := gin.Default()
	//在这里插入中间件
	r.Use(midware.RateLimitMiddleWare(time.Second, 1000, 1000))
	r.Use(midware.LogMidware())
	cs := config.NewControllerSet()
	//在这里插入controller
	cs.ADD("/regist", config.Method_POST, RegistUser)
	cs.ADD("/login", config.Method_POST, LoginUser)
	config.SetRouter(r, path, cs)
	return r
}
