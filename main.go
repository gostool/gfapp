package main

import (
	_ "gfapp/boot"
	_ "gfapp/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
