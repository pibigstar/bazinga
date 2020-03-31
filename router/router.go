package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/bazinga/app/api/websites"
)

func init() {
	s := g.Server()

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/websites", websites.List)
		group.ALL("/like", websites.LikeIt)
	})

	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Middleware()
	})
}
