package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nsplnpbjy/bbs/config"
	"github.com/nsplnpbjy/bbs/datamod"
	"github.com/nsplnpbjy/bbs/midware"
	"github.com/nsplnpbjy/bbs/utils"
)

func InitEngine(path string) *gin.Engine {
	r := gin.Default()
	//在这里插入中间件
	r.Use(midware.JWT())
	r.Use(midware.RateLimitMiddleWare(time.Second, 1000, 1000))
	r.Use(midware.LogMidware())
	cs := config.NewControllerSet()
	//在这里插入controller
	cs.ADD(config.RegistUrl, config.Method_POST, RegistUser)
	cs.ADD(config.LoginUrl, config.Method_POST, LoginUser)
	cs.ADD("/test", config.Method_POST, Test)
	config.SetRouter(r, path, cs)
	return r
}

// 测试用
func Test(c *gin.Context) {
	if c.PostForm("key") == "0" {
		user := datamod.User{Id: "0000", Username: "testname", Password: "testpassword", Regist_time: time.Now().Unix(), Ideas_id: nil, Comments_id: nil}
		token, _ := utils.GenerateToken(user.Username, user.Password)
		c.JSON(http.StatusOK, user.DePassword().SuccessReturner(token))
	} else {
		user := datamod.User{Id: "0000", Username: "testname", Password: "testpassword", Regist_time: time.Now().Unix(), Ideas_id: nil, Comments_id: nil}
		c.JSON(http.StatusOK, user.DePassword().FailReturner())
	}
}
