package redis

import (
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/util/gconv"
)

var MRedis = func() redis {
	return redis{
		gredis.New(config),
	}
}()

var (
	config = gredis.Config{
		Host: "127.0.0.1",
		Port: 6379,
		Db:   0,
	}
)

type redis struct {
	*gredis.Redis
}

func Set(key string, value interface{}) error {
	return MRedis.Set(key, value)
}

func Get(key string) (interface{}, error) {
	return MRedis.Get(key)
}

func GetString(key string) string {
	return MRedis.GetString(key)
}

func (r *redis) Set(key string, value interface{}) error {
	_, err := r.Do("SET", key, value)
	return err
}

func (r *redis) Get(key string) (interface{}, error) {
	return r.Do("GET", key)
}

func (r *redis) GetString(key string) string {
	 result, err := r.Do("GET", key)
	if err != nil {
		return ""
	}
	return gconv.String(result)
}