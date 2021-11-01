package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	a, b := make(chan int), make(chan int)

	go func() { // 接收端
		defer wg.Done()

		for {
			select {
			case x, ok := <-a:
				if ok == false { // 如果通道关闭,则设置为nil,阻塞
					a = nil
					break
				}
				println("a", x)
			case x, ok := <-b:
				if ok == false { // 如果通道关闭,则设置为nil,阻塞
					b = nil
					break
				}
				println("b", x)
			}

			if a == nil && b == nil { // 全部结束,退出循环
				return
			}
		}
	}()

	go func() { // 发送端 a
		defer wg.Done()
		defer close(a)

		for i := 0; i < 3; i++ {
			a <- i * 10
		}

	}()

	go func() { // 发送端 b
		defer wg.Done()
		defer close(b)

		for i := 0; i < 5; i++ {
			b <- i * 10
		}

	}()

	wg.Wait()
}
