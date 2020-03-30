package websites

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/bazinga/app/db"
)

// List all websites
func List(r *ghttp.Request) {
	r.Response.Writeln(db.MCategory.ListWebsites())
}
