package model

import "encoding/json"

type Post struct {
	Id       int      `gorm:"AUTO_INCREMENT,PRIMARY_KEY" json:"id"`
	PostName string   `json:"post_name"`
	AdvJson  string   `json:"-"`
	Ads      []string `gorm:"-" json:"adv_info"`
}

func (p *Post) AfterFind() (err error) {
	if p.AdvJson != "" {
		p.Ads = make([]string, 0)
		json.Unmarshal([]byte(p.AdvJson), &p.Ads)
	}
	return
}
