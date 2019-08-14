package repo

import (
	"adv/model"
	"github.com/jinzhu/gorm"
)

type Repositorier interface {
	Connection() *gorm.DB
	SetConnection(db *gorm.DB)
}
type AbstractRepo struct {
	conn *gorm.DB
}

func (repo *AbstractRepo) Connection() *gorm.DB {
	return repo.conn
}

func (repo *AbstractRepo) SetConnection(db *gorm.DB) {
	repo.conn = db
}

func (repo *AbstractRepo) getList(page int, pageSize int, out interface{}) (count int) {
	repo.Connection().Model(model.Post{}).Count(&count).Offset((page - 1) * pageSize).Limit(pageSize).Find(out)
	return
}
