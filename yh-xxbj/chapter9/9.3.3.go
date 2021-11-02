package main

import (
	"fmt"
	"time"
)

func init() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		fmt.Println("init: ", time.Now())
		time.Sleep(2 * time.Second)
	}()

	<-done
}

func main() {
	fmt.Println("main: ", time.Now())
}
