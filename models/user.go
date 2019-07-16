package models

import (
	// "github.com/jinzhu/gorm"
	"github.com/google/uuid"
	"strings"
	"math/big"
	"github.com/jinzhu/gorm"
	"encoding/json"
)

// User struct is full table of user
type User struct {
	UserID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	FirstName string  `json:"firstName"`
	LastName string `json:"lastName"`
	ProfileImage string `json:"profileImage"`
	Key string `json:"key"`
	IsLogin bool `json:"isLogin"`
}

// ExistUserByEmail checks if ther has a user with same email
func ExistUserByEmail(email string) (bool, error) {
	var user User
	db.AutoMigrate(&User{})
	err := db.Where("email = ? ", email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.UserID != "" {
		return true, nil
	}

	return false, nil
}

// AddUser will create a new user
func AddUser(email, password, firstName, lastName string) error  {
	id, _ := uuid.NewUUID()
	var i big.Int;
	i.SetString(strings.Replace(id.String(), "-", "", 4), 16)
	user := User{
		UserID: (i.String())[0:9],
		Email: email,
		FirstName: firstName,
		LastName: lastName,
		Password: EncodePassword(password),
		IsLogin: false,
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser()  {
	
}

func FindUserByEmail(email string) ([]byte, error) {
	var user User
	
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	p, _ := json.Marshal(user)

	return p, nil
}

func FindUser(u []byte) (interface{}, error) {
	var (
		users []User
		user User
	)
	_ = json.Unmarshal(u, &user)
	
	err := db.Where(user).Find(&users).Error
	if err != nil  && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	for index := 0; index < len(users); index++ {
		users[index].Password = ""
	}

	return users, nil
}