package controllers

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsocketManager struct {
	clients   map[string]*websocket.Conn
	broadcast chan response.NotifyResponse
	mutex     sync.Mutex
}

var wsManager = WebsocketManager{
	clients:   make(map[string]*websocket.Conn),
	broadcast: make(chan response.NotifyResponse),
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for simplicity; customize as needed
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (w *WebsocketManager) RegisterClient(userID string, client *websocket.Conn) {
	w.mutex.Lock()
	w.clients[userID] = client
	w.mutex.Unlock()
}

func (w *WebsocketManager) UnregisterClient(userID string) {
	w.mutex.Lock()
	if client, ok := w.clients[userID]; ok {
		client.Close()
		delete(w.clients, userID)
	}
	w.mutex.Unlock()
}

func (w *WebsocketManager) BroadcastMessage(userID string, notify response.NotifyResponse) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if client, ok := w.clients[userID]; ok {
		err := client.WriteJSON(notify)
		if err != nil {
			fmt.Printf("Failed to send message to user %s: %v\n", userID, err)
			w.UnregisterClient(userID)
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
	token, exsist := c.Get("current_user")
	if !exsist {
		conn.Close()
		return
	}
	currentResUser := token.(response.UserResponse)
	userID := currentResUser.ID

	wsManager.RegisterClient(userID, conn)
	defer wsManager.UnregisterClient(userID)

	conn.SetPingHandler(func(appData string) error {
		return nil
	})

	go func() {
		for {
			err := conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				break
			}
			// Wait for 30 seconds before sending the next ping
			time.Sleep(30 * time.Second)
		}
	}()

	for {
		var notify response.NotifyResponse
		err := conn.ReadJSON(&notify)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// Handle unexpected close error
				fmt.Printf("WebSocket error: %v\n", err)
			}
			break
		}

		wsManager.BroadcastMessage(userID, notify)
	}
}

func SendNotification(userID string, notify response.NotifyResponse) {
	wsManager.BroadcastMessage(userID, notify)
}
