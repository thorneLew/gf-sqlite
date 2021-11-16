package main

import (
	_ "project-build/boot"
	_ "project-build/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
