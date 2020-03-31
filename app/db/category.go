package db

var (
	MCategory Category
)

type Category struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Order       int        `json:"order"`
	Display     bool       `json:"display"`
	Description string     `json:"description"`
	Websites    []*Website `json:"websites"`
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
	var ids []int
	for _, result := range results {
		ids = append(ids, result.Id)
	}

	var websites []*Website
	err = db.From(MWebsite.name()).
		Where("display = ? AND category in (?)", true, ids).
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
