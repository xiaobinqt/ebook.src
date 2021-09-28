package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"unicode/utf8"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	//fmt.Println("service = ", service)
	conn, err := net.Dial("tcp", service)
	checkErr(err)

	/**
	备注：
	如果使用 HTTP/1.1 版本需要加上
	Connection: close不然会保持连接继续发送下一个请求，就会出现 400 错误
	还要加上 Host 头信息
	相连的 \r\n 中间不能有空格
	*/

	// 组成 GET 报文
	//message := fmt.Sprintf("GET /get HTTP/1.0\r\n\r\n\r\n")
	//message := fmt.Sprintf("GET /get HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n\r\n")

	// 组成 POST 报文

	// HTTP/1.0
	body := `{"name":"吴彦祖"}`
	message := fmt.Sprintf("POST / HTTP/1.0\r\nContent-Type: application/json;charset:utf-8\r\nContent-Length: %d\r\n\r\n%s", utf8.RuneCountInString(body), body)

	// HTTP/1.1
	//body := `{"name":"吴彦祖"}`
	//message := fmt.Sprintf("POST / HTTP/1.1\r\nHost: %s\r\nConnection: close\r\nContent-Type: application/json;charset:utf-8\r\nContent-Length: %d\r\n\r\n%s", service, utf8.RuneCountInString(body), body)
	//
	_, err = conn.Write([]byte(message))
	checkErr(err)

	result, err := readFull(conn)
	checkErr(err)

	fmt.Println(string(result))
	os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal err:%s ", err.Error())
		os.Exit(1)
	}
}

func readFull(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return result.Bytes(), nil
}
