package main

import (
	"crypto/tls"
	"io"
	"log"
)

func main() {
	conn, err := tls.Dial("tcp", "127.0.0.1:8092", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Fatalf("client dial:%s ", err.Error())
	}

	defer conn.Close()
	log.Println("client connected to", conn.RemoteAddr())

	state := conn.ConnectionState()
	log.Println("client handshake: ", state.HandshakeComplete)
	log.Println("client mutual: ", state.NegotiatedProtocolIsMutual)

	message := "Hello \n"
	n, err := io.WriteString(conn, message)
	if err != nil {
		log.Fatalf("client: write %s ", err.Error())
	}

	log.Printf("client wrote %q (%d bytes)", message, n)

	reply := make([]byte, 256)
	n, err = conn.Read(reply)
	if err != nil {
		log.Fatalf("conn read err %s ", err.Error())
	}
	log.Printf("client read %q (%d bytes)", string(reply[:n]), n)
	log.Print("client exiting")

}
