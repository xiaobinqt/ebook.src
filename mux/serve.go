package main

import (
	"fmt"
	"net"

	"ebook.src/mux/mux"
)

type TestHandle struct{}

func (h *TestHandle) Detect(header []byte) bool {
	fmt.Println("header: ", string(header))
	if string(header) != "wwwww" {
		return false
	}
	return true
}

func (h *TestHandle) Handle(conn net.Conn) error {
	//header, err := bufio.NewReader(conn).ReadString('\n')
	//if err != nil {
	//	fmt.Println("handle err", err.Error())
	//	return err
	//}
	//
	fmt.Println("Handle ..... ")

	_, err := conn.Write([]byte("success"))
	if err != nil {
		fmt.Println("conn write fail")
	}
	return nil
}

func main() {
	mux := mux.NewMux(nil)
	mux.SetHeaderReadFull(5)
	mux.RegisterHandler("test", &TestHandle{})

	err := mux.ListenAndServe(":8000")
	if err != nil {
		fmt.Println("listen and serve err", err.Error())
	}

}
