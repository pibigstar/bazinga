package redis

import (
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/util/gconv"
)

// TODO: 直接获取config的内容
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

func SetEX(key string, value interface{}, ex int) error {
	return MRedis.SetEX(key, value, ex)
}

func SetNX(key string, value interface{}) (bool, error) {
	result, err := MRedis.SetNX(key, value)
	if err != nil {
		return false, err
	}
	if result == "OK" {
		return true, nil
	}
	return false, nil
}

func SetNPX(key string, value interface{}, px int) (bool, error) {
	result, err := MRedis.SetNPX(key, value, px)
	if err != nil {
		return false, err
	}
	if result == "OK" {
		return true, nil
	}
	return false, nil
}

func Get(key string) (interface{}, error) {
	return MRedis.Get(key)
}

func GetString(key string) string {
	return MRedis.GetString(key)
}

func (r *redis) DoGet(args ...interface{}) error {
	_, err := r.Do("GET", args...)
	return err
}

func (r *redis) DoSet(args ...interface{}) error {
	_, err := r.Do("SET", args...)
	return err
}

func (r *redis) Set(key string, value interface{}) error {
	_, err := r.Do("SET", key, value)
	return err
}

func (r *redis) SetEX(key string, value interface{}, ex int) error {
	_, err := r.Do("SET", key, value, "EX", ex)
	return err
}

func (r *redis) SetPX(key string, value interface{}, px int) error {
	_, err := r.Do("SET", key, value, "PX", px)
	return err
}

func (r *redis) SetNX(key string, value interface{}) (interface{}, error) {
	return r.Do("SET", key, value, "NX")
}

func (r *redis) SetNPX(key string, value interface{}, px int) (interface{}, error) {
	return r.Do("SET", key, value, "PX", px, "NX")
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