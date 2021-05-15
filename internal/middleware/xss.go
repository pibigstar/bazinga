package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/bazinga/utils/xss"
	"log"
)

// xss过滤
func RemoveXss() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		err := xss.GetXss().XssRemove(r)
		if err != nil {
			log.Println(err)
		}
		r.Middleware.Next()
	}
}
