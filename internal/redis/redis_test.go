package redis

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

var (
	key = "bazinga"
	value = "test-redis"
	ex = 10
)

func TestSet(t *testing.T)  {
	err := Set(key, value)
	gtest.Assert(err, nil)

	err = SetEX(key, value, ex)
	gtest.Assert(err, nil)

	result, err := SetNX(key, value)
	gtest.Assert(err, nil)
	gtest.Assert(result, false)
}

func TestGetString(t *testing.T)  {
	s := GetString("bazinga")
	t.Log(s)
}