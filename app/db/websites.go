package db

var (
	MWebsites Websites
)

type Websites struct {
	Id          string `orm:"id"`
	Name        string `orm:"name"`
	Url         string `org:"url"`
	Category    int    `org:"category"`
	Order       int    `org:"order"`
	Display     bool   `org:"display"`
	Description string `org:"description"`
}

func (*Websites) name() string {
	return "websites"
}

func (w *Websites) List() (results []*Websites, err error) {
	err = db.From(w.name()).Where("id = ?", 1).Structs(&results)
	return results, err
}
