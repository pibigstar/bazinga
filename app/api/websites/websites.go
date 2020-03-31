package websites

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/bazinga/app/db"
	"github.com/pibigstar/bazinga/app/util/resp"
)

// List all websites
func List(r *ghttp.Request) {
	results, err := db.MCategory.ListWebsites()
	if err != nil {
		resp.Error(r, err.Error())
	}
	resp.SuccessWithDate(r, results)
}

func LikeIt(r *ghttp.Request) {
	// 检查IP是否已点过

	// 检查秘钥是否已存在

	err := db.MWebsite.LikeIt(r.GetString("id"))
	if err != nil {
		resp.Error(r, err.Error())
	}
	resp.WriteSuccess(r)
}
