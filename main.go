package main

import (
	"github.com/nsplnpbjy/bbs/config"
	"github.com/nsplnpbjy/bbs/service"
)

func main() {
	config.DbInit()
	r := service.InitEngine("/api")
	r.Run(":8092")
}
