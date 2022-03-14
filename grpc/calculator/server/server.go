package main

import (
	"context"
	"fmt"
	"net"

	"go.src/grpc/calculator/protobuf"
	"google.golang.org/grpc"
)

// 实现: CalculatorServiceServer接口, 在calculator.pb.go中定义
type server struct{}

func (s server) Calc(ctx context.Context, req *protobuf.CalcRequest) (resp *protobuf.CalcResponse, err error) {
	a := req.GetA()
	b := req.GetB()
	op := req.GetOp()
	resp = &protobuf.CalcResponse{}

	switch op {
	case "+":
		resp.R = a + b
	case "-":
		resp.R = a - b
	case "*":
		resp.R = a * b
	case "/":
		if b == 0 {
			err = fmt.Errorf("divided by zero")
			return
		}
		resp.R = a / b
	}
	return
}

// 启动rpc server
func main() {
	listener, err := net.Listen("tcp", "localhost:3233")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	protobuf.RegisterCalculatorServiceServer(s, &server{})
	fmt.Println("server start")
	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}
}
