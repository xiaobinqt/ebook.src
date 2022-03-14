package main

import (
	"context"
	"fmt"
	"log"

	"go.src/grpc/calculator/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连上grpc server
	//conn, err := grpc.Dial("localhost:3233", grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:3233", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := protobuf.NewCalculatorServiceClient(conn)

	// 调用远程方法
	resp, err := c.Calc(context.Background(), &protobuf.CalcRequest{
		A:  1,
		B:  2,
		Op: "+",
	})
	if err != nil {
		fmt.Println("calc err: ", err.Error())
		return
	}
	fmt.Println("calc success,respR: ", resp.GetR()) // 3
}
