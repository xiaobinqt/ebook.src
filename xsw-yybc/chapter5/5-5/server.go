package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"go.ebook.src/xsw-yybc/chapter5/5-5/server"
)

func main() {
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	http.Serve(l, nil)
}
