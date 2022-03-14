package main

import (
	"context"
	"fmt"
	"time"
)

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	go reqTask(ctx, "worker1")
	go reqTask(ctx, "worker2")

	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}

//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//
//	go reqTask(ctx, "worker1")
//	go reqTask(ctx, "worker2")
//
//	time.Sleep(3 * time.Second)
//	cancel()
//	time.Sleep(3 * time.Second)
//}
