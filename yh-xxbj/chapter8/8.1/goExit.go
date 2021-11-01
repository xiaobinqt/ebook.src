package main

import "runtime"

func main() {
	exit := make(chan struct{})

	go func() {
		defer close(exit)  // 执行
		defer println("a") // 执行

		func() {
			defer func() {
				// 执行 recover 返回 nil
				println("b", recover() == nil)
			}()

			func() { // 在多层调用中执行 goexit
				println("c")
				runtime.Goexit()   // 立即终止整个调用堆栈
				println("c done.") // 不会执行
			}()

			println("b done.") // 不会执行
		}()

		println("a done") // 不会执行
	}()

	<-exit
	println("main exit")

}
