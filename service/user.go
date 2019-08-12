package service

import (
	"adv/helper"
	"adv/model"
	"adv/repo"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	repo *repo.UserRepo
}

func (us *UserService) Login(username string, password string) (user *model.User, err error) {
	user = us.repo.GetUserByName(username)
	if user != nil && helper.CompareHashAndPassword([]byte(user.Password), []byte(password)) {
		return user, nil
	}
	return nil, fmt.Errorf("invalid username or bad password")
}

func NewUserService(db *gorm.DB) *UserService {
	us := new(UserService)
	us.repo = repo.NewUserRepo(db)
	return us
}
