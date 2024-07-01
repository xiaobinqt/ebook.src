package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
使用go实现1000个并发控制并设置执行超时时间1秒
*/

func worker(c context.Context, wg *sync.WaitGroup, id int) {

	defer wg.Done()

	select {
	case <-time.After(time.Second):
		fmt.Println("执行完成 %", id)
	case <-c.Done():
		fmt.Println("请求超时 %", id)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go worker(ctx, &wg, i)
	}
	wg.Wait()
}
