package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(1)

	exit := make(chan struct{})

	go func() {
		defer close(exit)

		go func() { // 任务 b, 放在此处为了确保 a 优先执行
			println("b")
		}()

		for i := 0; i < 4; i++ {
			println("a: ", i)
			if i == 1 { // 让出当前线程,调度执行 b
				runtime.Gosched()
			}
		}
	}()

	<-exit
}
