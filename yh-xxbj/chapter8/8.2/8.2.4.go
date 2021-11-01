package main

func main() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for x := range c {
			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c) // 如果不及时关闭会导致死锁..

	<-done
	println("done.....")
}
