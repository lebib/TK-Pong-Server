package main

import (
    "fmt"
    "Server/Game"
    "github.com/gorilla/websocket"
)

func main() {
    upgrader := websocket.Upgrader{
        ReadBufferSize: 1024,
        WriteBufferSize: 1024,
    }
    connection := Game.Websocket{upgrader}
    game := new(Game.Game)
    game.Init()
    fmt.Println(game.GameState2Json())
    fmt.Println(connection.Upgrader.ReadBufferSize)
}

// package main

// import (
//     "fmt"
//     "github.com/gorilla/websocket"
//     "net/http"
//     "Server/Game"
//     "Server/Web"
// )

// func TestMovePaddle() {
//     paddle := Game.Paddle{Game.Player{"Test", 0}, Game.Rect{0, 0, 100, 15}}
//     paddle.Move(100, Game.Rect{0, 0, 500, 500})
// }

// // func EchoHandler(w http.ResponseWriter, r *http.Request) {
// //     conn, err := upgrader.Upgrade(w, r, nil)
// //     if err != nil {
// //         return
// //     }

// //     for {
// //         messageType, p, err := conn.ReadMessage()
// //         if err != nil {
// //             return
// //         }

// //         PrintBinary(p)

// //         err = conn.WriteMessage(messageType, p);
// //         if err != nil {
// //             return
// //         }
// //     }
// // }

// func main() {
//     // http.HandleFunc("/echo", EchoHandler)
//     game := new(Game.Game)
//     game.Init
//     fmt.Printf("Hello!")
//     http.HandleFunc("/test", Web.SendMessage)
//     http.Handle("/", http.FileServer(http.Dir(".")))
//     err := http.ListenAndServe(":8080", nil)
//     if err != nil {
//         panic("Error: " + err.Error())
//     }
// }