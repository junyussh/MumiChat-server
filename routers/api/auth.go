package api

import (
	// "MumiChat/models"
	"MumiChat/pkg/app"
	"MumiChat/pkg/e"
	"MumiChat/routers/api/v1"
	"MumiChat/service/auth_service"
	"MumiChat/service/user_service"
	"fmt"
	"github.com/gorilla/websocket"
	"encoding/json"
)

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

type Auth struct {
	Email    string
	Password string
}

func GetAuth(conn *websocket.Conn, data map[string]string) {
	var (
		appW = app.Websocket{C: conn}
	)
	if data["email"] == "" || data["password"] == "" {
		appW.SocketResponse(e.ERROR_AUTH, nil)
		return
	}
	auth_service := auth_service.Auth{
		Email:    data["email"],
		Password: data["password"],
	}
	user_service := user_service.User{
		Email: data["email"],
	}
	isExist, err := auth_service.Check()
	if err != nil {
		fmt.Println(err)
		appW.SocketResponse(e.ERROR_AUTH_FAILED, err)
		return
	}
	if isExist {
		isNotLogin, err := auth_service.CheckLogin()
		if err != nil {
			appW.SocketResponse(e.ERROR_MULTIPLE_LOGIN_FAILED, err)
			return
		}
		// user not login
		if isNotLogin != false {
			auth_service.ChangeLoginStatus()
			appW.SocketResponse(e.LOGIN_SUCCESS, nil)

			p, _ := user_service.FindUserByEmail()
			var user User
			_ = json.Unmarshal(p, &user)
			// fmt.Println(user.UserID)
			client := e.Client{
				ID:         user.UserID,
				Email:		data["email"],
				ErrorCount: 0,
				Socket:     conn,
				Send:       make(chan []byte),
			}
			e.Clients[&client] = true
			e.Managers[user.UserID] = &client
			// fmt.Printf("email: %s, id: %s\n", e.Managers[user.UserID].Email, e.Managers[user.UserID].ID)
			go v1.WriteMessage(&client)
			// go v1.WriteBroadcast()

			return
		}

		// user has login
		appW.SocketResponse(e.ERROR_MULTIPLE_LOGIN, nil)
		return
	} else {
		// e.ErrorCount++
		appW.SocketResponse(e.ERROR_AUTH, nil)
		// if e.ErrorCount >= 3 {
		// 	conn.Close()
		// }
	}
}

func CleanAllStatus()  {
	auth_service.CleanAllStatus()
}
func Logout(conn *websocket.Conn) {
	for client := range e.Clients {
		if client.Socket == conn {
			auth_service := auth_service.Auth{
				Email: client.Email,
			}
			auth_service.ChangeLoginStatus()
			return
		}
	}
}
