package main

import (
	"github.com/nsplnpbjy/bbs/config"
)

func main() {
	r := config.InitEngine("/api")
	r.Run(":8092")
}
