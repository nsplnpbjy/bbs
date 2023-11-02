package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nsplnp/bbs/config"
)

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "helloworld"})
}

func main() {
	r := gin.Default()
	cs := config.NewControllerSet()
	cs.ADD("/h", config.Method_GET, HelloWorld)
	config.SetRouter(r, "/test", cs)
	r.Run(":8087")
}
