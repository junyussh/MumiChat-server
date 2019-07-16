package routers

import (
	"github.com/gorilla/websocket"
	"MumiChat/routers/api/v1"
	"MumiChat/routers/api"
)
// every is_login will be false
func CleanLoginStatus() {
	api.CleanAllStatus()
}

func Action(conn *websocket.Conn, action string, data map[string]string)  {
	switch action {
	case "login":
		api.GetAuth(conn, data)
		break
	case "register":
		v1.AddUser(conn, data)
		break
	case "query":
		v1.FindUser(conn, data)
	case "delete":
		v1.DeleteUser(conn, data)
		break
	}
}

func Logout(conn *websocket.Conn)  {
	api.Logout(conn)
}