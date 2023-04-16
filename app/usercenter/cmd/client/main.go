package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:1002/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		fmt.Println(string(message))
	}
}
