package main

import (
	"log"
	"fmt"
	"code.google.com/p/go.net/websocket"
)

func main () {
	origin := "http://localhost/"
	url := "ws://localhost:8080/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
    	log.Fatal(err)
	}

	var msg = make([]byte, 512)
	var n int

	if n, err = ws.Read(msg); err != nil {
    	log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
	
	_, err := ws.Close()
}

