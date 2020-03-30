package db

var (
	MCategory Category
)

type Category struct {
	Id          int    `orm:"id"`
	Name        string `orm:"name"`
	Order       int    `orm:"order"`
	Display     bool   `orm:"display"`
	Description string `orm:"description"`
	Websites    Websites
}

func (*Category) name() string {
	return "category"
}

func (c *Category) List() (results []*Category, err error) {
	err = db.From(c.name()).Where("display = ?", true).Structs(&results)
	return results, err
}
