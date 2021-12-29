package main

import (
	"log"
	"net"
)

func main() {
	c, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("Error to open TCP connection: %s", err)
	}
	defer c.Close()

	log.Printf("TCP session open")

	b := []byte("wwwwwt11111111111111111111111111111111111111111\n")
	_, err = c.Write(b)
	if err != nil {
		log.Fatalf("Error writing TCP session: %s", err)
	}

	for {
		d := make([]byte, 100)
		_, err = c.Read(d)
		if err != nil {
			log.Fatalf("Error reading TCP session: %s", err)
		}
		log.Printf("reading data from server: %s\n", string(d))
	}
}
