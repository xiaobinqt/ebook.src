package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan int), make(chan int)

	go func() { // 接收端
		defer wg.Done()

		for {
			var (
				name string
				x    int
				ok   bool
			)

			select {
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}

			if ok == false { // 如果任意通道关闭,则终止接收
				return
			}

			println(name, x) // 输出接收的数据类信息
		}
	}()

	go func() { // 发送端
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 10; i++ {
			select { // 随机选择发送 channel
			case a <- i:
			case b <- i * 10:

			}
		}

	}()

	wg.Wait()
}
