package middleware

import (
	"MumiChat/pkg/e"
	"MumiChat/pkg/app"
	"MumiChat/routers/api/v1"
	"github.com/gorilla/websocket"
	// "fmt"

)

func CheckIsLogin(conn *websocket.Conn, data map[string]string)  {
	for client := range e.Clients {
		if client.Socket == conn {
			v1.SendMessage(client, data)
			return
		}
	}

	appW := app.Websocket{C: conn}
	appW.SocketResponse(e.ERROR_UNAUTHORIZED, nil)
	return
}