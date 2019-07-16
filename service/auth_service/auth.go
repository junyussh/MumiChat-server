package auth_service

import (
	"MumiChat/models"
)

type Auth struct {
	UserID string
	Email string
	Password string
	IsLogin bool
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Email, a.Password)
}

func (a *Auth) CheckLogin() (bool, error) {
	return models.CheckLogin(a.Email)
}

func (a *Auth) ChangeLoginStatus()  {
	models.ChangeLoginStatus(a.Email)
}

func CleanAllStatus()  {
	models.CleanAllStatus()
}