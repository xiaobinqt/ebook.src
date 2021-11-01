package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	go func() {
		for {
			select {
			case <-time.After(5 * time.Second):
				fmt.Println("timeout...")
				os.Exit(0)
			}
		}
	}()

	go func() {
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				fmt.Println(time.Now())
			}
		}
	}()

	<-(chan struct{})(nil) // 直接用 nil channel 阻塞进程
}
