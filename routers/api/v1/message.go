package v1

import (
	"encoding/json"
	"time"
	// "strconv"
	"fmt"
	// "github.com/gorilla/websocket"
	"MumiChat/pkg/app"
	"MumiChat/pkg/e"
	// "MumiChat/service/user_service"
)

// Message is the struct of message
type Message struct {
	ID string `json:"id"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func SendMessage(client *e.Client, data map[string]string) {
	var appW = app.Websocket{C: client.Socket}
	if(data["type"] == "broadcast") {
		message, _ := json.Marshal(&Message{
			Sender: client.ID,
			Content: data["content"],
			CreatedAt: time.Now().UTC().Format(time.RFC3339),
		})
		e.Broadcast <- message
		return
	}
	// user isn't online
	if _, ok := e.Managers[data["recipient"]]; !ok {
		appW.SocketResponse(e.ERROR_USER_NOT_ONLINE, nil)
		return
	} else {
		message := &Message{
			Sender:    client.ID,
			Recipient: data["recipient"],
			Content:   data["content"],
			CreatedAt: time.Now().UTC().Format(time.RFC3339),
		}
		fmt.Printf("%s send to %s success\n", client.Email, e.Managers[data["recipient"]].Email)
		fmt.Printf("time: %s\n", time.Now().Format(time.RFC3339))
		p, _ := json.Marshal(message)

		recipient := e.Managers[data["recipient"]]
		recipient.Send <- p
		appW.SocketResponse(e.MESSAGE_SUCCESS, nil)
		return
	}
}

func WriteMessage(c *e.Client)  {
	var appW = app.Websocket{C: c.Socket}
	defer func() {
        c.Socket.Close()
    }()

    for {
        select {
        case message, ok := <-c.Send:
            if !ok {
                appW.SocketResponse(e.ERROR_CONNECTION_CLOSED, nil)
                return
            }

            appW.MessageResponse(message)
		case message, _ := <- e.Broadcast:
			for _, client := range e.Managers {
				fmt.Println("message: "+string(message))
				fmt.Println("email: "+client.Email)
				client.Send <- message
			}
		}
    }
}
func WriteBroadcast()  {
	for {
		select {
		case message, _ := <- e.Broadcast:
			for _, client := range e.Managers {
				fmt.Println("message: "+string(message))
				fmt.Println("email: "+client.Email)
				client.Send <- message
			}
		}
	}
}