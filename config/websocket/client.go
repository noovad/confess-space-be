package websocket

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client represents a connected websocket client
type Client struct {
	Hub      *Hub
	Conn     *websocket.Conn
	Send     chan Message
	Username string
	Email    string
	Channel  string
}

// ReadPump pumps messages from the websocket connection to the hub.
func (c *Client) ReadPump() {
	defer func() {
		log.Printf("Client %s disconnected from channel %s", c.Username, c.Channel)
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	log.Printf("Client %s started reading on channel %s", c.Username, c.Channel)

	for {
		var message Message
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			} else {
				log.Printf("Read error from client %s: %v", c.Username, err)
			}
			break
		}

		// Set message channel and sender
		message.Channel = c.Channel
		message.Sender = c.Username
		message.CreatedAt = time.Now()

		log.Printf("Received message from %s on channel %s: %+v", c.Username, c.Channel, message)

		c.Hub.Broadcast <- message
	}
}

// WritePump pumps messages from the hub to the websocket connection.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		log.Printf("Client %s stopped writing on channel %s", c.Username, c.Channel)
		ticker.Stop()
		c.Conn.Close()
	}()

	log.Printf("Client %s started writing on channel %s", c.Username, c.Channel)

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				log.Printf("Send channel closed for client %s", c.Username)
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			log.Printf("Sending message to %s on channel %s: %+v", c.Username, c.Channel, message)
			err := c.Conn.WriteJSON(message)
			if err != nil {
				log.Printf("Write error to client %s: %v", c.Username, err)
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("Ping error to client %s: %v", c.Username, err)
				return
			}
		}
	}
}
