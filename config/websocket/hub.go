package websocket

import (
	"go_confess_space-project/model"
	"log"
	"sync"
	"time"
)

type MessageType string

const (
	MessageTypeChat  MessageType = "chat"
	MessageTypeUsers MessageType = "users"
)

type Message struct {
	Type      MessageType         `json:"type"`
	ID        string              `json:"id,omitempty"`
	Message   model.Message       `json:"message,omitempty"`
	Channel   string              `json:"channel,omitempty"`
	CreatedAt time.Time           `json:"created_at,omitempty"`
	UsersData []map[string]string `json:"users_data,omitempty"`
}

type Hub struct {
	Users map[string]map[string]*Client

	Broadcast chan Message

	Register chan *Client

	Unregister chan *Client

	mu sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Users:      make(map[string]map[string]*Client),
		Broadcast:  make(chan Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			if _, ok := h.Users[client.Channel]; !ok {
				h.Users[client.Channel] = make(map[string]*Client)
			}
			log.Printf("Register: user %s joined channel %s", client.Username, client.Channel)
			h.Users[client.Channel][client.Username] = client
			h.mu.Unlock()

			go h.broadcastUserList(client.Channel)

		case client := <-h.Unregister:
			h.mu.Lock()
			if channelClients, ok := h.Users[client.Channel]; ok {
				if _, ok := channelClients[client.Username]; ok {
					delete(channelClients, client.Username)
					close(client.Send)
					log.Printf("Unregister: user %s left channel %s", client.Username, client.Channel)
					if len(channelClients) == 0 {
						delete(h.Users, client.Channel)
						log.Printf("Channel %s is now empty and removed", client.Channel)
					}
				}
			}
			h.mu.Unlock()

			go h.broadcastUserList(client.Channel)

		case message := <-h.Broadcast:
			h.mu.Lock()
			if channelClients, ok := h.Users[message.Channel]; ok {
				log.Printf("Broadcast: message to channel %s from %s", message.Channel, message.Message.User.Username)
				for _, client := range channelClients {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(channelClients, client.Username)
						log.Printf("Broadcast: closed send channel for user %s in channel %s", client.Username, message.Channel)
					}
				}
			}
			h.mu.Unlock()
		}
	}
}

func (h *Hub) broadcastUserList(channel string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	var users []map[string]string

	if channelClients, ok := h.Users[channel]; ok {
		for username, client := range channelClients {
			users = append(users, map[string]string{
				"username":    username,
				"name":        client.Name,
				"avatar_type": client.AvatarType,
				"created_at":  client.CreatedAt.GoString(),
			})
		}

		message := Message{
			Type:      MessageTypeUsers,
			UsersData: users,
			Channel:   channel,
			CreatedAt: time.Now(),
		}

		for _, client := range channelClients {
			select {
			case client.Send <- message:
			default:
			}
		}
	}
}
