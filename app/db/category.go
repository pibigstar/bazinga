package db

import (
	"context"
)

var (
	MWebsiteCategory WebsiteCategory
)

type WebsiteCategory struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Order       int        `json:"order"`
	Display     bool       `json:"display"`
	Description string     `json:"description"`
	Websites    []*Website `json:"websites"`
}

func (*WebsiteCategory) name() string {
	return "website_category"
}

func (c *WebsiteCategory) List() (results []*WebsiteCategory, err error) {
	err = db.Table(c.name()).Where("display = ?", true).Structs(&results)
	return results, err
}

func (c *WebsiteCategory) ListWebsites(ctx context.Context) (results []*WebsiteCategory, err error) {
	err = db.Table(c.name()).Ctx(ctx).Where("display = 1").Structs(&results)
	var ids []int
	for _, result := range results {
		ids = append(ids, result.Id)
	}

	var websites []*Website
	err = db.Ctx(ctx).Table(MWebsite.name()).
		Where("display = 1 AND category in (?)", ids).
		Structs(&websites)
	for _, result := range results {
		for _, website := range websites {
			if result.Id == website.Category {
				result.Websites = append(result.Websites, website)
			}
		}
	}

	return results, err
}
