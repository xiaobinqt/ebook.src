package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ready := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			println(id, ":ready") // 运动员准备就绪
			<-ready               // 等待发令
			println(id, "running....")
		}(i)
	}

	time.Sleep(time.Second)
	println("Ready? go! ")
	close(ready)

	wg.Wait()
}
