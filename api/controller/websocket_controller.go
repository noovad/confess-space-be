package controller

import (
	"fmt"
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

func NewWebSocketController(hub *websocket.Hub) *WebSocketController {
	go hub.Run()
	return &WebSocketController{hub: hub}
}

func (c *WebSocketController) HandleWebSocket(ctx *gin.Context) {
	username := ctx.Query("username")
	name := ctx.Query("name")
	avatarType := ctx.Query("avatar_type")
	channel := ctx.Query("channel")

	if username == "" || name == "" || avatarType == "" || channel == "" {
		err := fmt.Errorf("missing parameters: username, name, avatar_type, or channel")
		responsejson.BadRequest(ctx, err, "")
		return
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Printf("WebSocket upgrade error: %v\n", err)
		return
	}

	client := &websocket.Client{
		Hub:        c.hub,
		Conn:       conn,
		Send:       make(chan websocket.Message, 256),
		Username:   username,
		Name:       name,
		AvatarType: avatarType,
		Channel:    channel,
	}

	c.hub.Register <- client

	go func() {
		client.WritePump()
	}()
	go func() {
		client.ReadPump()
	}()
}
