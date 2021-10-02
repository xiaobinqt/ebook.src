package main

import (
	"crypto/rand"
	"crypto/tls"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("xsw.crt", "xsw.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s \n", err.Error())
	}

	config := tls.Config{
		Certificates: []tls.Certificate{
			cert,
		},
	}

	config.Time = time.Now
	config.Rand = rand.Reader

	service := "127.0.0.1:8092"

	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatalf("server: listen: %s ", err.Error())
	}

	log.Println("server: listening")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server:accept :%s ", err.Error())
			break
		}

		log.Printf("server : accepted from: %s ", conn.RemoteAddr())

		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)
	for {
		log.Print("server: conn: waiting")

		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("server conn read : %s ", err.Error())
			}
			break
		}

		log.Printf("server conn read echo %q \n", string(buf[:n]))
		n, err = conn.Write(buf[:n])
		log.Printf("server conn wrote :%d bytes ", n)

		if err != nil {
			log.Printf("server write : %s ", err.Error())
			break
		}
	}

	log.Println("server conn closed")
}
