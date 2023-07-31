package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Message struct {
	Type    string `json: "type"`
	Name    string `json: "name"`
	Message string `json: "message"`
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow connections from any origin (for testing purposes)
		},
	}

	clients []websocket.Conn
)

func Handler(c echo.Context) error {
	ws, _ := upgrader.Upgrade(c.Response(), c.Request(), nil)
	// defer ws.Close()

	clients = append(clients, *ws)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			return err
		}

		for _, client := range clients {
			err = client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				return err
			}
		}
	}
}
