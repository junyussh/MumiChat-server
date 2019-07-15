package app

import (
	"github.com/gorilla/websocket"

	"MumiChat/pkg/e"
	"encoding/json"
)

// Websocket struct
type Websocket struct {
	C *websocket.Conn
}

// Response format
type SocketResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Websocket) SocketResponse(errCode int, data interface{}) {
	response := SocketResponse{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	}
	b, _ := json.Marshal(response)
	g.C.WriteMessage(1, b)
	return
}
