package main

func main() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done) // 确保发出结束通知

		for {
			x, ok := <-c
			if ok == false { // 据此判断通道是否被关闭
				return
			}

			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)

	<-done
}
