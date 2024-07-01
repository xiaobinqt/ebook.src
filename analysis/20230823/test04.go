package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(15 * time.Second)
		ch <- struct{}{}
	}()
	<-ch

	fmt.Println(1111111)
}
