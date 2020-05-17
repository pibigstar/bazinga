package story

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/pibigstar/bazinga/app/db"
	"github.com/pibigstar/bazinga/app/util/resp"
)

func Add(r *ghttp.Request) {
	results, err := db.MWebsiteCategory.ListWebsites()
	if err != nil {
		resp.Error(r, err.Error())
	}
	resp.SuccessWithDate(r, results)
}

func Delete(r *ghttp.Request) {
	results, err := db.MWebsiteCategory.ListWebsites()
	if err != nil {
		resp.Error(r, err.Error())
	}
	resp.SuccessWithDate(r, results)
}

func Get(r *ghttp.Request) {
	results, err := db.MWebsiteCategory.ListWebsites()
	if err != nil {
		resp.Error(r, err.Error())
	}
	resp.SuccessWithDate(r, results)
}

func RandGet(r *ghttp.Request) {
	results, err := db.MWebsiteCategory.ListWebsites()
	if err != nil {
		resp.Error(r, err.Error())
	}
	resp.SuccessWithDate(r, results)
}

func Send(r *ghttp.Request) {
	results, err := db.MWebsiteCategory.ListWebsites()
	if err != nil {
		resp.Error(r, err.Error())
	}
	resp.SuccessWithDate(r, results)
}