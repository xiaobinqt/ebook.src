package main

import "fmt"

func test2(a ...int) {
	for i := range a {
		a[i] += 100
	}

	a = append(a, 2000)
	fmt.Println(a)
}

func main() {
	//a := []int{10, 20, 30}
	//test2(a...)

	a := [3]int{10, 20, 30}
	test2(a[:]...)

	fmt.Println(a)
}
