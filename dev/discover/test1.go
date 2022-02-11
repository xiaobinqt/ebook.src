package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//释放资源
	defer cancel()
	client := &http.Client{Transport: &http.Transport{}}
	resultChan := make(chan Result, 1)

	//发起请求
	//req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		fmt.Println("http request failed, err:", err)
		return
	}
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	pack := Result{r: resp, err: err}
	//将返回信息写入管道(正确或者错误的)
	resultChan <- pack

	select {
	case <-ctx.Done():
		er := <-resultChan
		fmt.Println("Timeout!", er.err)
		t, ok := ctx.Deadline()
		fmt.Println(t.Format("2006-01-02 15:04:05"), ok)
	case res := <-resultChan:
		fmt.Println("res == ", res)
		fmt.Println("res err == ", res.err)
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %d", len(out))
	}
	return
}

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	process()
}
