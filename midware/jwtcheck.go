package midware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nsplnpbjy/bbs/config"
	"github.com/nsplnpbjy/bbs/datamod"
	"github.com/nsplnpbjy/bbs/utils"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isSuccess := true
		var returnBody *datamod.Returner
		if config.AllowPathSet.Contains(ctx.Request.URL.Path) {
			ctx.Next()
		} else {
			token := ctx.PostForm("token")
			if token == "" {
				isSuccess = false
				returnBody = datamod.BlankTokenReturner()
			} else {
				claims, err := utils.ParseToken(token)
				if err != nil || claims == nil {
					isSuccess = false
					returnBody = datamod.InvalidTokenReturner()
				} else if time.Now().Unix() > claims.ExpiresAt {
					isSuccess = false
					returnBody = datamod.TimeOutTokenReturner()
				}
			}
		}
		if !isSuccess {
			ctx.JSON(http.StatusOK, returnBody)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
