package db

import (
	"github.com/gogf/gf/frame/g"
	_ "github.com/pibigstar/bazinga/internal/mysql"
)

var (
	db = g.DB()
)
