package db

import "time"

var (
	MWebsite Website
)

type Website struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	Category    int       `json:"category"`
	Order       int       `json:"order"`
	Score       int       `json:"score"`
	Display     bool      `json:"display"`
	CreateBy    string    `json:"create_by,omitempty"`
	CreateAt    time.Time `json:"create_at"`
	Description string    `json:"description"`
}

func (*Website) name() string {
	return "websites"
}

func (w *Website) List() (results []*Website, err error) {
	err = db.From(w.name()).Where("display = ?", false).Structs(&results)
	return results, err
}

func (w *Website) LikeIt(id string) error {
	_, err := db.From(w.name()).
		Where("id = ?", id).
		Update("score=score+1")
	return err
}
