package models

import (
	"encoding/base64"
	"golang.org/x/crypto/sha3"
	"github.com/jinzhu/gorm"
	// "fmt"
)

func EncodePassword(cleartext string) string {
	if cleartext == "" {
		return ""
	}
	h := make([]byte, 64)
	sha3.ShakeSum256(h, []byte(cleartext))
	return base64.StdEncoding.EncodeToString(h)
}

func CleanAllStatus()  {
	var user User
	db.Find(&user).Update("is_login", false)
}

func ChangeLoginStatus(email string)  {
	var user User
	db.Where(&User{
		Email: email,
	}).First(&user).Update("is_login", !user.IsLogin)
}
// false is login and true is not login
func CheckLogin(email string) (bool, error) {
	var user User
	err := db.Select("is_login").Where(&User{
		Email: email,
	}).First(&user).Error

	if err != nil {
		return false, err
	}
	if user.IsLogin != false {
		return false, nil
	}
	return true, nil
}

func CheckAuth(email, password string) (bool, error) {
	var auth User
	err := db.Where(User{
		Email: email,
		Password: EncodePassword(password),
	}).First(&auth).Error
	// fmt.Printf("email: %s, password: %s", auth.Email, auth.Password)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if auth.Email != "" {
		return true, nil
	}
	// not found user
	return false, nil
}