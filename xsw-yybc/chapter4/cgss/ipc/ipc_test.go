package ipc

import (
	"fmt"
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{
		Code: "ok",
		Body: "ECHO:" + method + " " + params,
	}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("foo", "From Client1")
	resp2, _ := client2.Call("foo", "From Client2")

	//fmt.Println("=======", resp1.Body)
	if resp1.Body != "ECHO:foo From Client1" || resp2.Body != "ECHO:foo From Client2" {
		fmt.Println("IpcClient.Call failed. resp1:", resp1.Body, "resp2:", resp2.Body)
		return
	}

	client1.Close()
	client2.Close()
	fmt.Println("------------ success -----------------")
}
