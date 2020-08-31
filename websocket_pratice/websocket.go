package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// websocket.Upgrader struct
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// loop to read any incomming msg
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))
		// if string(p) == "hello" || string(p) == "world" {
		// 	fmt.Println("Specific")
		// 	str := ""
		// 	fmt.Scan(&str)
		// 	conn.WriteMessage(1, []byte(str))
		// }
		if err = conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// wsEndPoint function for handling '/ws' route
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "wsEndpoint")

	// CheckOrigin is to determine whether allow to connect or not!
	// For now we just set it to true
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")

	// sender message to client
	err = ws.WriteMessage(1, []byte("Hi, Client!"))
	if err != nil {
		log.Println(err)
	}

	// reading msg from client
	reader(ws)
}

// homepage function for handling '/' route
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

// Setup the routes
func setuproutes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
  // fmt.Println("Hello World")

	// Setup the routes
	setuproutes()

	log.Fatal(http.ListenAndServe(":3000", nil))
}
