package midware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LogMidware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf(c.Request.RemoteAddr + "  访问IP")
		c.Next()
		log.Printf(c.Request.RemoteAddr + "  访问结束")
	}
}
