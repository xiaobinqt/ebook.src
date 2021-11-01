package main

const (
	max     = 50000000 // 数据统计上限
	block   = 500      // 数据块大小
	bufsize = 100      // 缓冲区大小
)

func Test() { // 普通模式,每次传递一个整数
	done := make(chan struct{})
	c := make(chan int, bufsize)

	go func() {
		count := 0
		for x := range c {
			count += x
		}
		close(done)
	}()

	for i := 0; i < max; i++ {
		c <- i
	}

	close(c)
	<-done

}

// 块模式, 每次将 500 个数字打包成块传输
func TestBlock() {
	done := make(chan struct{})
	c := make(chan [block]int, bufsize)

	go func() {
		count := 0
		for a := range c {
			for _, x := range a {
				count += x
			}
		}
		close(done)
	}()

	for i := 0; i < max; i += block {
		var b [block]int // 使用数组对数据打包
		for n := 0; n < block; n++ {
			b[n] = i + n
			if i+n == max-1 {
				break
			}
		}
		c <- b
	}

	close(c)
	<-done
}

func main() {

}
