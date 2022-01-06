package main

import "fmt"

func fo(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fo(n-1) + fo(n-2)
}

func main() {
	x := fo(40)
	fmt.Println(x)
}
