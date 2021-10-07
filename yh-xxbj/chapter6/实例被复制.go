package main

import "fmt"

type N int

func (n N) value() {
	//n++
	fmt.Printf("v:%p,%v \n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("p:%p,%v \n", n, *n)
}

func main() {
	var a N = 25

	a.pointer()
	a.value()

	fmt.Printf("a:%p,%v \n", &a, a)
}
