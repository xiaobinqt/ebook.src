package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

type Aab struct {
	Cpus []float64
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		time.Sleep(5 * time.Second)
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		fmt.Println("server read", string(msg))
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe("localhost:8011", nil))
}
