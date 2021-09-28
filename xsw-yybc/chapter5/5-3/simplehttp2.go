package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)

	fmt.Println("tcpAddr = ", tcpAddr)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErr(err)

	_, err = conn.Write([]byte("GET /hello HTTP/1.0\r\n\r\n"))
	checkErr(err)

	result, err := ioutil.ReadAll(conn)
	checkErr(err)

	fmt.Println("response:------------------- ")
	fmt.Println(string(result))

	os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal err:%s ", err.Error())
		os.Exit(1)
	}
}
