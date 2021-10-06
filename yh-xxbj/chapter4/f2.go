package main

import "fmt"

func test(x *int) {
	fmt.Printf("pointer: %p,target:%v \n", &x, x)
}

func main() {
	a := 0x100
	p := &a

	fmt.Printf("pointer: %p,target: %v \n", &p, p)

	test(p)
}
