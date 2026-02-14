package websocket

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域
	},
}

type Hub struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

// GlobalHub 全局 hub
var GlobalHub = &Hub{
	clients: make(map[*websocket.Conn]bool),
}

// ServerWS ws 接口
func (h *Hub) ServerWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	h.mu.Lock()
	h.clients[conn] = true
	h.mu.Unlock()

	// 保持连接
	go func() {
		defer func() {
			h.mu.Lock()
			delete(h.clients, conn)
			h.mu.Unlock()
			conn.Close()
		}()
		for {
			if _, _, err = conn.ReadMessage(); err != nil {
				break
			}
		}
	}()
}

// Broadcast 广播函数
func (h *Hub) Broadcast(message interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for client := range h.clients {
		_ = client.WriteJSON(message)
	}
}
