package e

import (
	"github.com/gorilla/websocket"
)

// store current user
var ( 
	User string
	ErrorCount = 0
	Clients = make(map[*Client]bool)
	Managers = make(map[*websocket.Conn]*Client)
)

type Client struct {
    ID     string
    ErrorCount int
    Socket *websocket.Conn
    Send   chan []byte
}