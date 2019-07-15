package user_service

import (
	"MumiChat/models"
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

func (u *User) Add() error {
	return models.AddUser(u.Email, u.Password, u.FirstName, u.LastName)
}

func (u *User) ExistByEmail() (bool, error) {
	return models.ExistUserByEmail(u.Email)
}

func (u *User) FindUserByEmail() ([]byte, error) {
	return models.FindUserByEmail(u.Email)
}