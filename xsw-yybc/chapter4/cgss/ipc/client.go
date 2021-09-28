package ipc

import (
	"encoding/json"
	"fmt"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()

	return &IpcClient{
		conn: c,
	}
}

func (client *IpcClient) Call(method, params string) (resp *Response,
	err error) {
	req := &Request{
		Method: method,
		Params: params,
	}

	var b = make([]byte, 0)
	b, err = json.Marshal(req)
	if err != nil {
		fmt.Println("call marshal err: ", err.Error())
		return nil, err
	}

	client.conn <- string(b)
	str := <-client.conn

	var resp1 Response
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1

	return resp, nil
}

func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
