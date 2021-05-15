package redis

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/gogf/gf/frame/g"
	_ "github.com/pibigstar/bazinga/boot"
)

var (
	client *redis.Client
	config *redis.Options
)

func InitRedis() {
	var cfg *redis.Options
	err := g.Cfg().GetStruct("redis", &cfg)
	if err != nil {
		panic(err)
	}
	config = cfg
	client = redis.NewClient(cfg)
}

func GetClient(ctx context.Context) *redis.Client {
	if client == nil {
		InitRedis()
	}
	cli := client.WithContext(ctx)
	cli.WrapProcess(func(oldProcess func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
		return func(cmd redis.Cmder) error {
			info := &callbackInfo{
				ctx:        ctx,
				attachment: cmd,
			}
			Before(info)
			defer After(info)
			return oldProcess(cmd)
		}
	})

	cli.WrapProcessPipeline(func(oldProcess func([]redis.Cmder) error) func([]redis.Cmder) error {
		return func(cmders []redis.Cmder) error {
			info := &callbackInfo{
				ctx:        ctx,
				attachment: cmders,
			}
			Before(info)
			defer After(info)
			return oldProcess(cmders)
		}
	})

	return cli
}
