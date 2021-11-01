package main

import "time"

var c int

func counter() int {
	c++
	return c
}

func main() {
	a := 100

	go func(x, y int) {
		//x1, y1 := x, y
		time.Sleep(time.Second)
		//println("go: ", x1, y1)
		println("go: ", x, y)
	}(a, counter())

	a += 100

	println("main: ", a, counter())
	time.Sleep(3 * time.Second)

}
