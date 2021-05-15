package db

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/pibigstar/bazinga/utils/seq"
	"time"
)

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
	CreateBy    int       `json:"create_by,omitempty"`
	CreateName  string    `json:"create_name,omitempty"`
	Description string    `json:"description,omitempty"`
	CreateAt    time.Time `json:"create_at"`
}

func (*Website) name() string {
	return "websites"
}

func (w *Website) List() (results []*Website, err error) {
	err = db.Table(w.name()).Where("display = ?", false).Structs(&results)
	return results, err
}

func (w *Website) LikeIt(id string) (string, error) {
	_, err := db.Table(w.name()).
		Where("id = ?", id).
		Update("score=score+1")
	if err != nil {
		return "", err
	}

	// 生成个uid
	uid := seq.GenID()
	fmt.Println("uid", uid)
	_, err = g.Redis().Do("SET", uid, "ex", time.Hour*24*3)
	if err != nil {
		return "", err
	}
	return uid, err
}
