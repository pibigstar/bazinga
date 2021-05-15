package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/pibigstar/bazinga/internal/grpc/server"
	"github.com/pibigstar/bazinga/internal/middleware"

	_ "github.com/pibigstar/bazinga/boot"
	_ "github.com/pibigstar/bazinga/router"
)

func main() {
	s := g.Server()
	// 中间件
	middleware.Init(s)
	// 启动grpc
	go server.StartGrpc()

	s.Run()
}
