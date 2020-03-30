package db

import "github.com/gogf/gf/os/glog"

var (
	MCategory Category
)

type Category struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Order       int         `json:"order"`
	Display     bool        `json:"display"`
	Description string      `json:"description"`
	Websites    []*Websites `json:"websites"`
}

func (*Category) name() string {
	return "category"
}

func (c *Category) List() (results []*Category, err error) {
	err = db.From(c.name()).Where("display = ?", true).Structs(&results)
	return results, err
}

func (c *Category) ListWebsites() (results []*Category, err error) {
	err = db.From(c.name()).Where("display = ?", true).Structs(&results)

	for _, result := range results {
		err = db.From(MWebsites.name()).Where("display = ? AND category = ?", true, result.Id).Structs(&result.Websites)
		if err != nil {
			glog.Printf("scan err: %v", err)
			return nil, err
		}
	}
	return results, err
}
