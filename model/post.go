package model

type Post struct {
	Id       int    `gorm:"AUTO_INCREMENT,PRIMARY_KEY" json:"id"`
	PostName string `json:"post_name"`
	AdvJson  string `json:"adv_info"`
}
