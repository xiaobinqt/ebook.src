package main

import (
	"fmt"
	"net/rpc"

	"go.ebook.src/xsw-yybc/chapter5/5-5/server"
)

func main() {
	//rpc.ClientCodec()
	//rpc.ServerCodec()

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("dial http err", err.Error())
		return
	}

	args := &server.Args{
		A: 7,
		B: 8,
	}

	// 同步调用
	//var reply int
	//err = client.Call("Arith.Multiply", args, &reply)
	//if err != nil {
	//	log.Fatal("arith err", err)
	//}
	//
	//fmt.Printf("arith: %d*%d=%d", args.A, args.B, reply)

	// 异步调用
	quotient := new(server.Quotient)
	divCall := client.Go("Arith.Divide", args, &quotient, nil)
	//replyCall := <-divCall.Done
	<-divCall.Done

	//fmt.Println(replyCall, quotient)
	fmt.Println(quotient)
}
