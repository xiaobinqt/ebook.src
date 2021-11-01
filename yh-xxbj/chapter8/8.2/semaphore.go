package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	sem := make(chan struct{}, 2) // 最多允许 2 个并发执行

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			sem <- struct{}{} // 获取信号

			defer func() { // 释放信号
				<-sem
			}()

			time.Sleep(2 * time.Second)
			fmt.Println(id, time.Now())

		}(i)
	}

	wg.Wait()
}
