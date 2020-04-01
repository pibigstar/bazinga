package redis

import (
	"testing"
)

func TestSet(t *testing.T)  {
	err := Set("bazinga", "test-redis")
	if err != nil {
		panic(err)
	}
}

func TestGetString(t *testing.T)  {
	s := GetString("bazinga")
	t.Log(s)
}