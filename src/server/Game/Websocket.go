package Game

import (
    "github.com/gorilla/websocket"
    "net/http"
    "fmt"
)

////////////////////////////////////////////////////////////////////////////////

type Websocket struct {
    websocket.Upgrader
}

////////////////////////////////////////////////////////////////////////////////

func (wSocket *Websocket) SendMessage(w http.ResponseWriter, r *http.Request, m string) {
    conn, err := wSocket.Upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    err = conn.WriteMessage(1, []byte(m))
}

func (wSocket *Websocket) ReceiveMessage(w http.ResponseWriter, r *http.Request) {
    conn, err := wSocket.Upgrader.Upgrade(w, r, nil)
    messageType, p, err := conn.ReadMessage()
    if err != nil {
        return
    }
    fmt.Printf(string(messageType))
    fmt.Printf(string(p))
}

////////////////////////////////////////////////////////////////////////////////