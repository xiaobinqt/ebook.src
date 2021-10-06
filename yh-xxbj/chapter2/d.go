package main

import "fmt"

func add(a, _ int) int {
	fmt.Println(a)
	return a
}

func main() {
	println(add(1, 2))
}
