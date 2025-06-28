package controller

import (
	"log"
	"net/http"

	"go_confess_space-project/config/websocket"
	"go_confess_space-project/helper/responsejson"

	"github.com/gin-gonic/gin"
	gorilla "github.com/gorilla/websocket"
)

var upgrader = gorilla.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketController struct {
	hub *websocket.Hub
}

func NewWebSocketController() *WebSocketController {
	hub := websocket.NewHub()
	go hub.Run()
	log.Println("[WebSocket] Hub started and running")
	return &WebSocketController{hub: hub}
}

func (c *WebSocketController) HandleWebSocket(ctx *gin.Context) {
	username := ctx.Query("username")
	email := ctx.Query("email")
	channel := ctx.Query("channel")

	if username == "" || email == "" || channel == "" {
		responsejson.BadRequest(ctx, nil, "Missing required query parameters: username, email, channel")
		return
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("[WebSocket] Upgrade error: %v", err)
		return
	}
	log.Printf("[WebSocket] Connection established: username=%s, email=%s, channel=%s", username, email, channel)

	client := &websocket.Client{
		Hub:      c.hub,
		Conn:     conn,
		Send:     make(chan websocket.Message, 256),
		Username: username,
		Email:    email,
		Channel:  channel,
	}

	c.hub.Register <- client
	log.Printf("[WebSocket] Client registered: username=%s, channel=%s", username, channel)

	go func() {
		log.Printf("[WebSocket] WritePump started for username=%s", username)
		client.WritePump()
	}()
	go func() {
		log.Printf("[WebSocket] ReadPump started for username=%s", username)
		client.ReadPump()
	}()
}

func (c *WebSocketController) SendMessage(message websocket.Message) {
	log.Printf("[WebSocket] Broadcasting message to channel=%s", message.Channel)
	c.hub.Broadcast <- message
}
