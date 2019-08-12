package repo

import (
	"adv/model"
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	AbstractRepo
}

func (repo *UserRepo) GetUserByName(username string) *model.User {
	user := new(model.User)
	if repo.Connection().Where("username = ?", username).First(user).Error == nil {
		return user
	}
	return nil
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	repo := new(UserRepo)
	repo.SetConnection(db)
	return repo
}
