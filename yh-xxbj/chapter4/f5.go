package main

func test3() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func main() {
	x := test3()(1, 2)
	println(x)
}
