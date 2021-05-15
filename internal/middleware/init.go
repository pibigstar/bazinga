package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Init(s *ghttp.Server) {

	s.Use(Timestamp())

	if g.Cfg().GetBool("rate.apiRate") {
		s.Use(ApiLate())
	}

	if g.Cfg().GetBool("rate.ipRate") {
		s.Use(IpLate())
	}

	if g.Cfg().GetBool("trace.enable") {
		s.Use(Trace())
	}

	s.Use(RemoveXss(), Metrics())
}
