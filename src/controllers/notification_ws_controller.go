package controllers

import (
	"net/http"
	"sync"

	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsocketManager struct {
	clients   map[*websocket.Conn]bool
	broadcast chan response.NotifyResponse
	mutex     sync.Mutex
}

var wsManager = WebsocketManager{
	clients:   make(map[*websocket.Conn]bool),
	broadcast: make(chan response.NotifyResponse),
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for simplicity; customize as needed
		return true
	},
}

func (w *WebsocketManager) RegisterClient(client *websocket.Conn) {
	w.mutex.Lock()
	w.clients[client] = true
	w.mutex.Unlock()
}

func (w *WebsocketManager) UnregisterClient(client *websocket.Conn) {
	w.mutex.Lock()
	delete(w.clients, client)
	w.mutex.Unlock()
}

func (w *WebsocketManager) BroadcastMessage(notify response.NotifyResponse) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	for client := range w.clients {
		err := client.WriteJSON(notify)
		if err != nil {
			w.UnregisterClient(client)
			client.Close()
		}
	}
}

func HandleWebsocketConnection(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade to WebSocket"})
		return
	}
	defer conn.Close()

	wsManager.RegisterClient(conn)
	defer wsManager.UnregisterClient(conn)

	for {
		var notify response.NotifyResponse
		err := conn.ReadJSON(&notify)
		if err != nil {
			break
		}

		wsManager.BroadcastMessage(notify)
	}
}

func SendNotification(notify response.NotifyResponse) {
	wsManager.BroadcastMessage(notify)
}
