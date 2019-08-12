package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password []byte) []byte {
	ret, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return ret
}

func CompareHashAndPassword(hash []byte, password []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, password) == nil
}
