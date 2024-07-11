package config

import "github.com/gin-gonic/gin"

type MethodType string

type controllerInfo struct {
	path       string
	method     MethodType
	handleFunc func(*gin.Context)
}

type ControllerSet struct {
	set map[int]controllerInfo
}

func NewControllerSet() *ControllerSet {
	controllerset := ControllerSet{make(map[int]controllerInfo)}
	return &controllerset
}

func (c *ControllerSet) ADD(p string, m MethodType, f func(*gin.Context)) {
	controllerinfo := controllerInfo{
		path:       p,
		method:     m,
		handleFunc: f,
	}
	c.set[len(c.set)] = controllerinfo
}
