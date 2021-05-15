package websites

import (
	"encoding/json"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/pibigstar/bazinga/app/db"
	"github.com/pibigstar/bazinga/internal/consts"
	"github.com/pibigstar/bazinga/internal/redis"
	"github.com/pibigstar/bazinga/utils/resp"
	"time"
)

// List all websites
func List(r *ghttp.Request) {
	var results []*db.WebsiteCategory
	rds := redis.GetClient(r.Context())
	res, err := rds.Get(consts.RedisKeyWebsites).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(res), &results); err != nil {
			resp.Error(r, err)
			return
		}
		resp.Success(r, results)
		return
	}

	results, err = db.MWebsiteCategory.ListWebsites(r.Context())
	if err != nil {
		resp.Error(r, err)
		return
	}
	rds.Set(consts.RedisKeyWebsites, gconv.String(results), time.Hour*1)

	resp.Success(r, results)
}

func LikeIt(r *ghttp.Request) {
	// 检查IP是否已点过

	// 检查秘钥是否已存在
	_, err := db.MWebsite.LikeIt(r.GetString("id"))
	if err != nil {
		resp.Error(r, err)
	}
	resp.Success(r, nil)
}
