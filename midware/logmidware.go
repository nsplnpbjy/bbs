package midware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LogMidware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("  访问IP:" + c.Request.RemoteAddr)
		c.Next()
		log.Printf("  访问结束" + c.Request.RemoteAddr)
	}
}
