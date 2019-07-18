package e

import (
	"github.com/gorilla/websocket"
)

// store current user
var ( 
	// ErrorCount = 0
	Clients = make(map[*Client]bool)
	Managers = make(map[string]*Client)
)

type Client struct {
	ID     string
	Email string
    ErrorCount int
    Socket *websocket.Conn
    Send   chan []byte
}