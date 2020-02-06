package repo

import (
	"adv/model"
	"github.com/jinzhu/gorm"
)

// repo is base repository with common methods
type repo struct {
	conn *gorm.DB
}

func (repo *repo) Connection() *gorm.DB {
	return repo.conn
}

func (repo *repo) SetConnection(db *gorm.DB) {
	repo.conn = db
}

func (repo *repo) getList(page int, pageSize int, out interface{}) (count int) {
	repo.Connection().Model(model.Post{}).Count(&count).Offset((page - 1) * pageSize).Limit(pageSize).Find(out)
	return
}
