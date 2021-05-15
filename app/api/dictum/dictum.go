package dictum

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/bazinga/utils/resp"
	"time"
)

func Random(r *ghttp.Request) {
	time.Sleep(1 * time.Second)
	resp.Success(r, "Hello")
}
