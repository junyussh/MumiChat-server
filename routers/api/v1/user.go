package v1

import (
	"github.com/gorilla/websocket"
	"MumiChat/pkg/app"
	"MumiChat/service/user_service"
	"MumiChat/pkg/e"
)

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

// AddUserForm is a form of register
type AddUserForm struct {
	Email string
	FirstName string 
	LastName string
}

// AddUser register a user
func AddUser(conn *websocket.Conn, data map[string]string)  {
	var appW = app.Websocket{C: conn}
	user_service := user_service.User{
		Email: data["email"],
		FirstName: data["firstName"],
		LastName: data["lastName"],
		Password: data["password"],
	}
	exists, err := user_service.ExistByEmail()
	if err != nil {
		appW.SocketResponse(e.ERROR_EXIST_USER_FAILED, err)
		return
	}
	if exists {
		appW.SocketResponse(e.ERROR_EXIST_USER, nil)
		return
	}
	err = user_service.Add()
	if err != nil {
		appW.SocketResponse(e.ERROR_ADD_USER_FAILED, nil)
		return
	}
	appW.SocketResponse(e.SUCCESS, nil)
}

// DeleteUser will delete a user
func DeleteUser(conn *websocket.Conn, data map[string]string)  {
	
}
