package models

import (
	// "github.com/jinzhu/gorm"
	"github.com/google/uuid"
	"strings"
	"math/big"
	"github.com/jinzhu/gorm"
	// "fmt"
	"encoding/json"
)

// User struct is full table of user
type User struct {
	UserID string 
	Username string 
	Email string
	Password string
	FirstName string 
	LastName string
	ProfileImage string
	Key string
	IsLogin bool
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
	// fmt.Println(p)
	// fmt.Printf("%+v", user)
	return p, nil
}