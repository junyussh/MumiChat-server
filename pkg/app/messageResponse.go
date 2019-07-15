package app

type MessageResponse struct {
	Sender string
	Recipient string
	Content string
	CreatedAt string
}

func (g *Websocket) MessageResponse(p []byte) {
	g.C.WriteMessage(1, p)
	return
}