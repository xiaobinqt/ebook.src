package main

import "time"

func main() {
	exit := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		println("go routine done")

		time.Sleep(10 * time.Second)
		close(exit) // 关闭通道,发出信号
	}()

	println("main...")
	<-exit // 如通道关闭，立即解除阻塞
	println("main exit")
}
