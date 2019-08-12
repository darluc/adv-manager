package model

type User struct {
	Id       int `gorm:"AUTO_INCREMENT,PRIMARY_KEY"`
	Username string
	Password string
	avatar   string
	name     string
}
