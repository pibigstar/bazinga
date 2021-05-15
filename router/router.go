package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/bazinga/app/api/dictum"
	"github.com/pibigstar/bazinga/app/api/websites"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	s := g.Server()

	// 注册全局router
	globalRouter(s)

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/websites", websites.List)
		group.ALL("/like", websites.LikeIt)

	})

	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Middleware()
	})

	s.Group("/dictum", func(group *ghttp.RouterGroup) {
		group.GET("random", dictum.Random)
	})
}

func globalRouter(s *ghttp.Server) {
	handler := promhttp.Handler()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/metrics", func(r *ghttp.Request) {
			handler.ServeHTTP(r.Response.Writer, r.Request)
		})
	})
}
