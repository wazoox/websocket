package main

import (
	"net/http"
	"log"
	"code.google.com/p/go.net/websocket"
	"time"
	"math"
	"strconv"
)

func main() {
	log.Println("Starting Server...")

	http.Handle("/ws", websocket.Handler(handler));

	err := http.ListenAndServe("127.0.0.1:8080", nil);

	if err != nil {
		log.Fatal("ListenAndServe: ", err )
	}
}

// global var: all clients see the same curve
var x float64 = 0


func handler(ws *websocket.Conn) {
	
	for {
		if x >= 2*math.Pi {
			x = 0
		} else {
			x += 0.05
		}
		
		time.Sleep(500*1000*1000) // sleep for 500ms (Sleep takes nanoseconds)

		msg := strconv.FormatFloat(math.Sin(x), 'g', 10, 64)
		log.Printf("%v sending: %v\n", ws, msg)
		websocket.Message.Send( ws, msg )
	}
}

