package repo

import (
	"adv/model"
	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	GetUserByName(username string) *model.User
}

type userRepo struct {
	repo
}

func (repo *userRepo) GetUserByName(username string) *model.User {
	user := new(model.User)
	if repo.Connection().Where("username = ?", username).First(user).Error == nil {
		return user
	}
	return nil
}

func NewUserRepo(db *gorm.DB) UserRepo {
	repo := new(userRepo)
	repo.SetConnection(db)
	return repo
}
