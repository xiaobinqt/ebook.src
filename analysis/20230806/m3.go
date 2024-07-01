package main

import "fmt"

func accumulator() func(int) int {
	var x int

	return func(delta int) int {
		fmt.Printf("(%+v,%+v)-", &x, x)
		x += delta
		return x
	}
}

func main() {
	var a = accumulator()
	fmt.Printf("%d \n", a(1))
	fmt.Printf("%d \n", a(10))
	fmt.Printf("%d \n", a(100))

	fmt.Println("----------")

	var b = accumulator()
	fmt.Printf("%d \n", b(1))
	fmt.Printf("%d \n", b(10))
	fmt.Printf("%d \n", b(100))
}
