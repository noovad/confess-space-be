package websocket

import (
	"log"
	"sync"
	"time"
)

// MessageType defines different types of messages
type MessageType string

const (
	MessageTypeChat  MessageType = "chat"
	MessageTypeUsers MessageType = "users"
)

// Message represents a message sent over WebSocket
type Message struct {
	Type      MessageType         `json:"type"`
	ID        string              `json:"id,omitempty"`
	Content   string              `json:"content,omitempty"`
	Sender    string              `json:"sender,omitempty"`
	Receiver  string              `json:"receiver,omitempty"`
	Channel   string              `json:"channel,omitempty"`
	CreatedAt time.Time           `json:"created_at,omitempty"`
	UsersData []map[string]string `json:"users_data,omitempty"`
}

// Hub maintains active clients and broadcasts messages
type Hub struct {
	// Registered clients by channel
	Users map[string]map[string]*Client // map[channel][username]*Client

	// Inbound messages from clients
	Broadcast chan Message

	// Register requests from clients
	Register chan *Client

	// Unregister requests from clients
	Unregister chan *Client

	mu sync.Mutex
}

// NewHub creates a new hub instance
func NewHub() *Hub {
	return &Hub{
		Users:      make(map[string]map[string]*Client),
		Broadcast:  make(chan Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// Run starts the hub to handle client connections and messages
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			if _, ok := h.Users[client.Channel]; !ok {
				h.Users[client.Channel] = make(map[string]*Client)
			}
			h.Users[client.Channel][client.Username] = client
			h.mu.Unlock()

			log.Printf("[REGISTER] User '%s' joined channel '%s'", client.Username, client.Channel)

			go h.broadcastUserList(client.Channel)

		case client := <-h.Unregister:
			h.mu.Lock()
			if channelClients, ok := h.Users[client.Channel]; ok {
				if _, ok := channelClients[client.Username]; ok {
					delete(channelClients, client.Username)
					close(client.Send)
					log.Printf("[UNREGISTER] User '%s' left channel '%s'", client.Username, client.Channel)
					if len(channelClients) == 0 {
						delete(h.Users, client.Channel)
						log.Printf("[CHANNEL EMPTY] Channel '%s' deleted", client.Channel)
					}
				}
			}
			h.mu.Unlock()

			go h.broadcastUserList(client.Channel)

		case message := <-h.Broadcast:
			h.mu.Lock()
			if channelClients, ok := h.Users[message.Channel]; ok {
				for _, client := range channelClients {
					select {
					case client.Send <- message:
						log.Printf("[BROADCAST] Message sent to '%s' in channel '%s'", client.Username, message.Channel)
					default:
						close(client.Send)
						delete(channelClients, client.Username)
						log.Printf("[ERROR] Failed to send message to '%s' in channel '%s', client unregistered", client.Username, message.Channel)
					}
				}
			}
			h.mu.Unlock()
		}
	}
}

// broadcastUserList sends the current list of active users to all clients in a channel
func (h *Hub) broadcastUserList(channel string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	var users []map[string]string

	if channelClients, ok := h.Users[channel]; ok {
		for username, client := range channelClients {
			users = append(users, map[string]string{
				"username": username,
				"email":    client.Email,
			})
		}

		message := Message{
			Type:      MessageTypeUsers,
			UsersData: users,
			Channel:   channel,
		}

		for _, client := range channelClients {
			select {
			case client.Send <- message:
				log.Printf("[USERLIST] Sent user list to '%s' in channel '%s'", client.Username, channel)
			default:
				log.Printf("[ERROR] Failed to send user list to '%s' in channel '%s'", client.Username, channel)
			}
		}
	}
}
