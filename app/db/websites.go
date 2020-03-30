package db

var (
	MWebsites Websites
)

type Websites struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Category    int    `json:"category"`
	Order       int    `json:"order"`
	Display     bool   `json:"display"`
	Description string `json:"description"`
}

func (*Websites) name() string {
	return "websites"
}

func (w *Websites) List() (results []*Websites, err error) {
	err = db.From(w.name()).Where("display = ?", false).Structs(&results)
	return results, err
}
