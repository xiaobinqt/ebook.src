package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		wg.Wait()
		println("wait exit")
	}()

	go func() {
		time.Sleep(5 * time.Second)
		println("done")
		wg.Done()
	}()

	wg.Wait()
	println("main exit")

}
