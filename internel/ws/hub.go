package ws

import (
	"sync"
)

type Hub struct {
	clients    sync.Map
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	clientNum  int64
}

func NewHub() *Hub {
	return &Hub{
		clients:    sync.Map{},
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clientNum:  0,
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients.Store(client.userID, client)
		case client := <-h.unregister:
			if _, ok := h.clients.LoadAndDelete(client.userID); ok {
				close(client.send)
			}
		}
	}
}
