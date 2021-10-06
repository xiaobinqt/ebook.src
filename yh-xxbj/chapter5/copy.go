package main

import "fmt"

func test(x *[2]int) {
	fmt.Printf("x :%p,%v \n", x, *x)
	x[1] += 100
}

func main() {
	a := [2]int{10, 20}
	test(&a)

	fmt.Printf("a: %p,%v \n", &a, a)
}
