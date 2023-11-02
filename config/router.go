package config

import "github.com/gin-gonic/gin"

func SetRouter(r *gin.Engine, groupPath string, c *ControllerSet) {
	group := r.Group(groupPath)
	for i := 0; i <= len(c.set)-1; i++ {
		switch c.set[i].method {
		case Method_GET:
			group.GET(c.set[i].path, c.set[i].handleFunc)
		case Method_POST:
			group.POST(c.set[i].path, c.set[i].handleFunc)
		case Method_DELETE:
			group.DELETE(c.set[i].path, c.set[i].handleFunc)
		}
	}
}
