package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

/**
	消费

生产，用 协程往 channel 写数据

消费，消费 channel 里的数据


*/

var (
	dataChan = make(chan int, 0)
)

func produce(data int) {
	dataChan <- data
}

/**
熔断算法

*/

type Count struct {
	Count int
	Err   int
}

func consumer() {
	for {
		select {
		case d := <-dataChan:
			surl := "http://127.0.0.1:8088"
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, surl,
				strings.NewReader(fmt.Sprintf("%d", d)))

			fmt.Println("消费的数据为：", d, req, err)
			//case := <- notify:

		}
	}
}

func main() {
	go func() {
		for {
			produce(time.Now().Second())
			//time.Sleep(5 * time.Second)
		}
	}()

	go consumer()

	for {

	}
}
