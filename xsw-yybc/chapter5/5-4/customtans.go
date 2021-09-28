package main

import (
	"fmt"
	"net/http"
)

type OurCustomTransport struct {
	Transport http.RoundTripper
}

func (t *OurCustomTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *OurCustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// 处理一些事情
	// 发情 http 请求
	// 添加一些域到 req.Header 中
	req.Header.Add("X-h", "1111")
	fmt.Println("222")

	return t.transport().RoundTrip(req)
}

func (t *OurCustomTransport) Client() *http.Client {
	return &http.Client{
		Transport: t,
	}
}

func main() {
	t := &OurCustomTransport{}

	c := t.Client()
	resp, err := c.Get("http://127.0.0.1:8080/hello")
	if err != nil {
		fmt.Println("err = ", err.Error())
		return
	}

	fmt.Printf("resp = %+v \n", resp)
}
