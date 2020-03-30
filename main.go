package main

import (
	"github.com/gogf/gf/frame/g"
	_ "github.com/pibigstar/bazinga/boot"
	_ "github.com/pibigstar/bazinga/router"
)

func main() {
	g.Server().Run()
}
