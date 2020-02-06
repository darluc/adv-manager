package service

import (
	"adv/helper"
	"adv/model"
	"adv/repo"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserService interface {
	Login(username string, password string) (user *model.User, err error)
}

type userService struct {
	repo repo.UserRepo
}

func (us *userService) Login(username string, password string) (user *model.User, err error) {
	user = us.repo.GetUserByName(username)
	if user != nil && helper.CompareHashAndPassword([]byte(user.Password), []byte(password)) {
		return user, nil
	}
	return nil, fmt.Errorf("invalid username or bad password")
}

func NewUserService(db *gorm.DB) UserService {
	service := new(userService)
	service.repo = repo.NewUserRepo(db)
	return service
}
